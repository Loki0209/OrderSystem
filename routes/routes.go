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

		// Store routes
		stores := v1.Group("/stores")
		{
			// Public endpoints (for QR code scanning and customer viewing)
			stores.GET("/:id", controllers.GetStore)
			stores.GET("", controllers.GetAllStores)

			// Protected endpoints (require authentication - for store owners)
			storesProtected := stores.Group("")
			storesProtected.Use(middleware.AuthMiddleware())
			{
				storesProtected.POST("", controllers.CreateStore)
				storesProtected.GET("/my-stores", controllers.GetMyStores)
				storesProtected.PUT("/:id", controllers.UpdateStore)
				storesProtected.DELETE("/:id", controllers.DeleteStore)
				storesProtected.PATCH("/:id/toggle-status", controllers.ToggleStoreStatus)
			}
		}

		// Category routes
		categories := v1.Group("/categories")
		{
			// Public endpoints (for customers to view menu)
			categories.GET("/store/:storeId", controllers.GetCategoriesByStore)
			categories.GET("/store/:storeId/active", controllers.GetActiveCategoriesByStore)
			categories.GET("/:id", controllers.GetCategory)

			// Protected endpoints (require authentication - for store owners)
			categoriesProtected := categories.Group("")
			categoriesProtected.Use(middleware.AuthMiddleware())
			{
				categoriesProtected.POST("", controllers.CreateCategory)
				categoriesProtected.PUT("/:id", controllers.UpdateCategory)
				categoriesProtected.DELETE("/:id", controllers.DeleteCategory)
			}
		}

		// Food Item routes
		foodItems := v1.Group("/food-items")
		{
			// Public endpoints (for customers to view menu)
			foodItems.GET("/store/:storeId", controllers.GetFoodItemsByStore)
			foodItems.GET("/store/:storeId/available", controllers.GetAvailableFoodItemsByStore)
			foodItems.GET("/category/:categoryId", controllers.GetFoodItemsByCategory)
			foodItems.GET("/category/:categoryId/available", controllers.GetAvailableFoodItemsByCategory)
			foodItems.GET("/:id", controllers.GetFoodItem)

			// Protected endpoints (require authentication - for store owners)
			foodItemsProtected := foodItems.Group("")
			foodItemsProtected.Use(middleware.AuthMiddleware())
			{
				foodItemsProtected.POST("", controllers.CreateFoodItem)
				foodItemsProtected.PUT("/:id", controllers.UpdateFoodItem)
				foodItemsProtected.DELETE("/:id", controllers.DeleteFoodItem)
				foodItemsProtected.PATCH("/:id/toggle-availability", controllers.ToggleFoodItemAvailability)
			}
		}
	}

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Restaurant Ordering System API",
			"version": "1.0.0",
			"endpoints": gin.H{
				"health":      "/api/v1/hello",
				"register":    "POST /api/v1/auth/register",
				"login":       "POST /api/v1/auth/login",
				"users":       "/api/v1/users (requires auth)",
				"products":    "/api/v1/products (requires auth)",
				"stores":      "/api/v1/stores",
				"categories":  "/api/v1/categories",
				"food_items":  "/api/v1/food-items",
			},
		})
	})
}
