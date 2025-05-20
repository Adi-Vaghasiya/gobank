package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", store)

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	server := NewAPIserver(":"+port, store)
	server.Run()
}
