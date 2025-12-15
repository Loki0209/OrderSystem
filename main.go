package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"ordernew/config"
	"ordernew/routes"
	"ordernew/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Set Gin mode
	gin.SetMode(config.AppConfig.GinMode)

	// Connect to MongoDB
	err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer config.DisconnectDatabase()

	// Initialize collections
	services.InitUserCollection()
	services.InitProductCollection()
	services.InitStoreCollection()
	services.InitCategoryCollection()
	services.InitFoodItemCollection()

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Graceful shutdown
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		log.Println("Shutting down server...")
		config.DisconnectDatabase()
		os.Exit(0)
	}()

	// Start server
	serverAddress := ":" + config.AppConfig.Port
	log.Printf("Server starting on port %s...", config.AppConfig.Port)
	log.Printf("API Documentation available at http://localhost:%s/api/v1/hello", config.AppConfig.Port)

	if err := router.Run(serverAddress); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
