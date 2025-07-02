package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	store, err := NewDatabaseConn()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	fmt.Println(port)
	newServer := NewApi(":"+port, store)
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("your server is starting on: %v\n", port)
	newServer.Run()
}
