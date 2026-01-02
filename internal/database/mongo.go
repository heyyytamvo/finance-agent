package database

import (
	"context"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect(uri string, auth_mechanism string, user string, password string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

  credential := options.Credential{
//   		AuthMechanism: auth_mechanism,
  		Username:      user,
  		Password:      password,
  	}
  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  clientOpts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI).
  		SetAuth(credential)

  client, err := mongo.Connect(clientOpts)
  if err != nil {
    log.Panic(err)
  }


	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
