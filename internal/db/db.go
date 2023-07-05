package db

import (
	"fmt"
	"os"

	"github.com/Galbeyte1/gRPC-microservices/internal/rocket"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db * sqlx.DB
}


// New - returns a new store or error
func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_PORT")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbTable,
		dbPassword,
		dbSSLMode,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Store{}, err
	}
	return Store{
		db: db,
	}, nil
}

// GetRocketByID - returns a rocket from the database by a given ID
func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

// InsertRocket - inserts a new rocket into the database
func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

// DeleteRocket - deletes a rocket from the database by it's ID
func (s Store) DeleteRocket(id string) error {
	return nil
}