package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client = ConnectDB()

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(EnvMongoURI()),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("содинение с mongodb установлено...")
	return client
}

// TODO: имя базы в .env
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("application").Collection(collectionName)
	return collection
}
