package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	MongoURI    string
	DatabaseName string
	JWTSecret   string
	JWTExpiry   string
	GinMode     string
	APIVersion  string
}

var AppConfig *Config

// LoadConfig loads environment variables and initializes the application configuration
func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	AppConfig = &Config{
		Port:        getEnv("PORT", "8080"),
		MongoURI:    getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		DatabaseName: getEnv("DATABASE_NAME", "ordernew_db"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiry:   getEnv("JWT_EXPIRY", "24h"),
		GinMode:     getEnv("GIN_MODE", "debug"),
		APIVersion:  getEnv("API_VERSION", "v1"),
	}

	log.Println("Configuration loaded successfully")
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
