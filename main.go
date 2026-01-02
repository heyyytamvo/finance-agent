package main

import (
	"log"

	"my-finance-app/internal/config"
	"my-finance-app/internal/database"
	"my-finance-app/internal/server"
)

func main() {
	cfg := config.Load()

	client, err := database.Connect(cfg.MongoURI, cfg.AuthMechanism, cfg.Username, cfg.Password)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(cfg.MongoDB)

	srv := server.New(db)
  log.Fatal(srv.Run(":8080"))
}
