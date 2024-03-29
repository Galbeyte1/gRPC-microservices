package grpc

import (
	"context"
	"log"
	"net"

	"github.com/Galbeyte1/gRPC-microservices/internal/rocket"
	rkt "github.com/Galbeyte1/rocket-protos/rocket/v1"
	"google.golang.org/grpc"
)

// RocketService - define the interface that the concrete implementation
// has to adhere to
type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

// Handler - will handle incoming gRPC requests
type Handler struct{
	RocketService RocketService
}

// New - returns a new gRPC handler
func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)
	log.Println("Listening on 0.0.0.0:50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s\n", err)
		return err
	}
	return nil
}

// GetRocket - retrieves a rocket by ID and returns a response
func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	log.Print("Get Rocket gRPC Endpoint Hit")

	rocket, err := h.RocketService.GetRocketByID(ctx, req.Id)
	if err != nil {
		log.Println("Failed to retrieve rocket by ID")
		return &rkt.GetRocketResponse{}, err
	}

	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id: rocket.ID,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	return &rkt.AddRocketResponse{}, nil
}

func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	return &rkt.DeleteRocketResponse{}, nil
}