package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

func main() {
	startup()
	_ = provider.NewDatabase()
	fmt.Println("Should start server")
}
