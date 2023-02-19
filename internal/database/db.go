package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoURI() (string, error) {
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUsername := os.Getenv("DB_USERNAME")
	if dbUsername == "" || dbPassword == "" {
		return "", errors.New("ENV NOT FOUND")
	}
	return "mongodb+srv://" + dbUsername + ":" + dbPassword + "@cluster0.ncxolee.mongodb.net/?retryWrites=true&w=majority", nil
}

var BookCollection *mongo.Collection
var AuthorCollection *mongo.Collection

func ConnectDB() {
	dbName := os.Getenv("DB_NAME")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	connectionURI, err := getMongoURI()
	if err != nil {
		log.Fatal(err)
	}
	clientOptions := options.Client().
		ApplyURI(connectionURI).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	BookCollection = client.Database(dbName).Collection("Books")
	AuthorCollection = client.Database(dbName).Collection("Authors")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected successfully")

}
