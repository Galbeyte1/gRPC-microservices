package main

import (
	"fmt"
	"log"

	"github.com/Galbeyte1/gRPC-microservices/internal/db"
	"github.com/Galbeyte1/gRPC-microservices/internal/rocket"
)

func Run() error {
	// responsible for initializing and starting
	// our gRPC server
	fmt.Println("Rocket Service Starting...")
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	_ = rocket.New(rocketStore)
	return nil
}

func main() {
	// logging errors
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}