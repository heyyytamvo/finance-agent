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
  MongoDB     string
}

var (
	cfg  *Config
	once sync.Once
)

func Load() *Config {

  once.Do(func() {
  		err := godotenv.Load()
  		if err != nil {
  			log.Fatal("❌ .env file NOT loaded")
  		}
  		log.Println("✅ .env file loaded")


  		cfg = &Config{
  			MongoURI:       os.Getenv("MONGO_URI"),
//   			AuthMechanism:  os.Getenv("AUTHENMECHANISM"),
  			Username:       os.Getenv("MONGO_ADMIN_USER"),
  			Password:       os.Getenv("MONGO_ADMIN_PASSWORD"),
  			MongoDB:        os.Getenv("MONGO_DB"),
  		}
  	})
  	return cfg
}
