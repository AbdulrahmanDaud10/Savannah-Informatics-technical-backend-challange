package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AbdulrahmanDaud10/savannah-info-customer-order-service/pkg/route"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	return nil
}

func main() {
	fmt.Println("Customer Order Service Start")

	// Loading Environmental Variable file
	if err := LoadEnv(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to load .env file: %s\n", err)
		os.Exit(1)
	}

	if err := route.SetupRoutes(":8090"); err != nil {
		fmt.Fprintf(os.Stderr, "Server start failure: %s\n", err)
		os.Exit(1)
	}
}
