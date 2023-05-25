package main

import (
	"fmt"
	"log"
)

func Run() error {
	// responsible for initializing and starting
	// our gRPC server
	fmt.Println("Rocket Service Starting...")
	return nil
}

func main() {
	// logging errors
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}