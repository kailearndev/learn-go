package app

import (
	"log"
	"os"

	"kai-shop-be/internal/domain/category"
	"kai-shop-be/internal/domain/product"
	"kai-shop-be/internal/domain/upload"
	"kai-shop-be/internal/domain/user"
	"kai-shop-be/internal/server"
	"kai-shop-be/pkg/cloudstorage"
	"kai-shop-be/pkg/db"
)

// Run is the main entry point for the backend app
func Run() {
	// 1. Init Database
	database := db.InitPostgres()
	database.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`) // chỉ cần 1 lần mỗi DB
	// 2. AutoMigrate models
	if err := database.AutoMigrate(
		&product.Product{},
		&user.User{},
		&category.Category{},
	); err != nil {
		log.Fatalln("❌ migrate failed:", err)
	}
	categoryRepo := category.NewRepository(database)
	categoryService := category.NewService(categoryRepo)
	categoryHandler := category.NewHandler(categoryService)
	// 3. Init Repositories & Services
	productRepo := product.NewRepository(database)
	productService := product.NewService(productRepo, categoryRepo)
	productHandler := product.NewHandler(productService)
	//// user
	userRepo := user.NewRepository(database)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	// category

	// 4. Init CloudFly Storage
	storage, err := cloudstorage.NewCloudFlyConfig(
		os.Getenv("CLOUDFLY_ENDPOINT"), // e.g. "https://s3.cloudfly.vn"
		os.Getenv("CLOUDFLY_ACCESS_KEY"),
		os.Getenv("CLOUDFLY_SECRET_KEY"),
		os.Getenv("CLOUDFLY_BUCKET"), // e.g. "kai-shop-bucket"
	)
	if err != nil {
		log.Fatalln("❌ failed to connect CloudFly:", err)
	}

	// 5. Init Upload handler
	uploadHandler := upload.NewHandler(storage)

	// 6. Setup router with all handlers
	router := server.SetupRouter(
		productHandler,
		userHandler,
		uploadHandler,
		categoryHandler,
	)

	// 7. Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running at http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalln("❌ failed to start server:", err)
	}
}
