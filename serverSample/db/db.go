package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBHandler interface {
	InsertInDatabase(*Payload) error
}

type GlobalPayload struct {
	Payload
}

type Mongo struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

type DataEntry struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (db *Mongo) InitialConnection() (*Mongo, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	MONGO_URI := os.Getenv("MONGO_URI")
	log.Printf("Mongo URI: %s", MONGO_URI)
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		fmt.Println("Connected to mongoDB!!!")
	}
	collection := client.Database("sampleServer").Collection("sampleServer")

	return &Mongo{Client: client, Collection: collection}, nil
}

// InsertInDatabase implements MongoDbHandler interface.
func (db *Mongo) InsertInDatabase(payload *Payload) error {
	fmt.Println("The Payload", payload.Email, payload.Password)
	Pload := DataEntry{
		Email:    payload.Email,
		Password: payload.Password,
	}
	inserted, err := db.Collection.InsertOne(context.Background(), &Pload)
	if err != nil {
		return err
	}
	log.Printf("USer Inserted: %v", inserted.InsertedID)
	return nil
}
