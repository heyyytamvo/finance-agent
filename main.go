package main

import (
	"log"

	"my-finance-app/internal/config"
	"my-finance-app/internal/database"
	"my-finance-app/internal/server"
  "my-finance-app/internal/services/spending"
)

func main() {
	cfg := config.Load()

	// Connect to MongoDB
  client, err := database.Connect(cfg.MongoURI, cfg.AuthMechanism, cfg.Username, cfg.Password)
  if err != nil {
    log.Fatal(err)
  }
  db := client.Database(cfg.MongoDatabase)

  // Initialize repository & service
  spendingRepo := spending.NewRepository(db)
  spendingService := &spending.Service{Repo: spendingRepo}

  // Initialize server with the service
  srv := server.New(spendingService)

  log.Println("Server running on :8080")
  log.Fatal(srv.Run(":8080"))
}
