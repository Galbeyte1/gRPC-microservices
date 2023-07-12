//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/Galbeyte1/gRPC-microservices/main/internal/rocket Store

package rocket

import "context"

// Primary Service for providing rocket data

// Rocket - should contain the definition of our rocket
type Rocket struct {

	ID string 
	Name string
	Type string
	Flights int 

}

// Store - defines the interface we expect
// our database implementation to follow
type Store interface {
	// Store so far is housing Rocket() control
	// methods

	// Given Rocket object/ID

	// Identify a rocket in the store
	GetRocketByID(id string) (Rocket, error)
	// Add a new rocket to the store
	InsertRocket(rkt Rocket) (Rocket, error)
	// Remove rocket from the store
	DeleteRocket(id string) error
}


// Service - our rocket service responsible for
// updating the rocket inventory
type Service struct{

	Store Store
}


func New(store Store) Service {
	return Service{
		Store: store,
	}
}


// GetRocketByID - retrieves a rocket based on the ID from the strore
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil

}

// InsertRocket - inserts a rocket based on the Rocket object
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket - deletes a rocket from our inventory
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}  