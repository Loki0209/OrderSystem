package routes

import (
	"ordernew/controllers"
	"ordernew/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine) {
	// Initialize controllers
	userController := controllers.NewUserController()
	productController := controllers.NewProductController()

	// Apply CORS middleware
	router.Use(middleware.CORSMiddleware())

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		// Health check / Hello World endpoint
		v1.GET("/hello", userController.HelloWorld)

		// Public routes (Authentication)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", userController.Register)
			auth.POST("/login", userController.Login)
		}

		// Protected routes (require authentication)
		users := v1.Group("/users")
		users.Use(middleware.AuthMiddleware())
		{
			users.GET("", userController.GetAllUsers)
			users.GET("/:id", userController.GetUserByID)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}

		// Product routes (require authentication)
		products := v1.Group("/products")
		products.Use(middleware.AuthMiddleware())
		{
			products.POST("", productController.CreateProduct)                   // Create product
			products.GET("", productController.GetAllProducts)                   // Get all products
			products.GET("/search", productController.SearchProducts)            // Search products
			products.GET("/category/:category", productController.GetProductsByCategory) // Get by category
			products.GET("/:id", productController.GetProductByID)               // Get product by ID
			products.PUT("/:id", productController.UpdateProduct)                // Update product (full)
			products.PATCH("/:id", productController.PatchProduct)               // Patch product (partial)
			products.DELETE("/:id", productController.DeleteProduct)             // Delete product
			products.PUT("/:id/quantity", productController.UpdateProductQuantity) // Update quantity only
		}
	}

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Order Management API",
			"version": "1.0.0",
			"endpoints": gin.H{
				"health":   "/api/v1/hello",
				"register": "POST /api/v1/auth/register",
				"login":    "POST /api/v1/auth/login",
				"users":    "/api/v1/users (requires auth)",
				"products": "/api/v1/products (requires auth)",
			},
		})
	})
}
