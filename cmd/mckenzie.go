package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"mckenzie/provider"
)

func startup() {
	log.Println("Loading .env")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

// ! Follow returning structs and accepting interfaces!
func main() {
	startup()
	database := provider.NewDatabase()

	fmt.Println("Should start server")
}
