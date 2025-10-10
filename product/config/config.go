package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUsername string
	MongoPassword string 
	MongoHost string
	MongoPort string
	MongoDB  string
}

func LoadConfig() *Config {
	// Load .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		MongoUsername: getEnv("DB_USER", "mongodb://localhost:27017"),
		MongoPassword: getEnv("DB_PASSWORD", "mongodb://localhost:27017"),
		MongoHost: getEnv("DB_HOST", "mongodb://localhost:27017"),
		MongoPort: getEnv("DB_PORT", "mongodb://localhost:27017"),
		MongoDB:  getEnv("DB_NAME", "products"),
	}
}

// Helper for default values
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
