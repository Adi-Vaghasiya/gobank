package main

import (
	"context"
	"fmt"
	"log"
	db "serverSample/db"

	// "serverSample/db"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	d := &db.Mongo{}
	dbInstance, err := d.InitialConnection()
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer func() {
		if err := dbInstance.Client.Disconnect(ctx); err != nil {
			log.Fatalf("Mongo disconnect error: %v", err)
		}
		fmt.Println("MongoDB connection closed.")
	}()
	p := db.Payload{Mongo: db.Mongo{
		Client:     dbInstance.Client,
		Collection: dbInstance.Collection,
	}}
	p.Handler()

}
