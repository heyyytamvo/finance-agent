package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	AuthMechanism  string
	Username     string
  Password      string
  MongoDatabase     string
}

var (
	cfg  *Config
	once sync.Once
)

func Load() *Config {
	once.Do(func() {
		ginMode := os.Getenv("GIN_MODE")

		if ginMode != "release" {
			// Load .env ONLY in non-release mode
			if err := godotenv.Load(); err != nil {
				log.Println("⚠️ .env file not loaded (non-release mode)")
			} else {
				log.Println("✅ .env file loaded")
			}
		} else {
			log.Println("🚀 Running in RELEASE mode, using environment variables")
		}

		cfg = &Config{
			MongoURI:      os.Getenv("MONGO_URI"),
			Username:      os.Getenv("MONGO_ADMIN_USER"),
			Password:      os.Getenv("MONGO_ADMIN_PASSWORD"),
			MongoDatabase: os.Getenv("MONGO_DATABASE"),
		}
	})

	return cfg
}

