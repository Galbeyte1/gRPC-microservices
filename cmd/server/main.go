package main

import (
	"log"

	"github.com/Galbeyte1/gRPC-microservices/internal/db"
	"github.com/Galbeyte1/gRPC-microservices/internal/rocket"
)

func Run() error {
	// responsible for initializing and starting
	// our gRPC server
	log.Println("Rocket Service Starting...")
	rocketStore, err := db.New()
	if err != nil {
		log.Println("Failed to instantiate Database")
		return err
	}
	log.Println("Instantiated Database")

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}
	_ = rocket.New(rocketStore)
	log.Println("Migrations complete")
	return nil
}

func main() {
	// logging errors
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}