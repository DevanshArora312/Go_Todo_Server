package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var Client *mongo.Client

func ConnectDB() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	MONGO_URI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	Client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Cannot connect to MongoDB!")
	}
	if err = Client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}
	Collection = Client.Database("go-test").Collection("todos")
	fmt.Println("Connected to DB!")

}
