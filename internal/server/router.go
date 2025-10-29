package server

import (
	"kai-shop-be/internal/domain/category"
	"kai-shop-be/internal/domain/product"
	"kai-shop-be/internal/domain/upload"
	"kai-shop-be/internal/domain/user"
	"kai-shop-be/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	productHandler *product.Handler,
	userHandler *user.Handler,
	uploadHandler *upload.Handler,
	categoryHandler *category.Handler,
) *gin.Engine {
	r := gin.Default()

	// ======== PUBLIC ROUTES =========
	public := r.Group("/")
	{
		// Auth
		public.POST("/auth/register", userHandler.RegisterUser)
		public.POST("/auth/login", userHandler.LoginUser)

		// Products (public read)

	}

	// ======== PROTECTED ROUTES =========
	protected := r.Group("/")
	protected.Use(middleware.JWTAuth()) // middleware JWT check
	{
		// Product management (must login)
		protected.POST("/products", productHandler.CreateProduct)
		protected.PUT("/products/:id", productHandler.UpdateProduct)
		protected.DELETE("/products/:id", productHandler.DeleteProduct)
		protected.GET("/products/:id", productHandler.GetProductByID)
		protected.GET("/products", productHandler.ListProducts)
		protected.POST("/users", userHandler.RegisterUser)
		protected.PUT("/users/:id", userHandler.UpdateUser)
		protected.DELETE("/users/:id", userHandler.DeleteUser)
		protected.GET("/users/:id", userHandler.GetUserByID)
		protected.GET("/users", userHandler.ListUsers)
		// Category management
		protected.POST("/categories", categoryHandler.CreateCategory)
		protected.PUT("/categories/:id", categoryHandler.UpdateCategory)
		protected.DELETE("/categories/:id", categoryHandler.DeleteCategory)
		protected.GET("/categories/:id", categoryHandler.GetCategoryByID)
		protected.GET("/categories", categoryHandler.ListCategories)
		// Example: current user

		protected.POST("/upload", uploadHandler.UploadImage)
	}

	return r
}
