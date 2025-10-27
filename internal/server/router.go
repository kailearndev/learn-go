package server

import (
	"kai-shop-be/internal/domain/product"
	"kai-shop-be/internal/domain/upload"
	"kai-shop-be/internal/domain/user"

	"github.com/gin-gonic/gin"
)

// RouteConfig gom các handler lại để dễ truyền vào SetupRouter
type RouteConfig struct {
	ProductHandler *product.Handler
	UploadHandler  *upload.Handler
	UserHandler    *user.Handler
	// sau này có thể thêm:
	// CategoryHandler *category.Handler
	// AuthHandler     *auth.Handler
}

// SetupRouter initializes all routes
func SetupRouter(cfg RouteConfig) *gin.Engine {
	r := gin.Default()

	// Register domain routes
	if cfg.ProductHandler != nil {
		cfg.ProductHandler.RegisterRoutes(r)
	}

	if cfg.UploadHandler != nil {
		cfg.UploadHandler.RegisterRoutes(r)
	}

	if cfg.UserHandler != nil {
		cfg.UserHandler.RegisterRoutes(r)
	}
	// future:
	// if cfg.CategoryHandler != nil {
	//     cfg.CategoryHandler.RegisterRoutes(r)
	// }

	return r
}
