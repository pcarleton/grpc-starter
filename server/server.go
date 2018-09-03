package server

import (
	pb "github.com/pcarleton/grpc-starter/proto/api"
	"golang.org/x/net/context"
)

type server struct {
}

func NewServer() pb.ApiServer {
	return &server{}
}

func (s *server) GetHealth(ctx context.Context, request *pb.GetHealthRequest) (*pb.GetHealthResponse, error) {
	return &pb.GetHealthResponse{
		Status: pb.HealthStatus_OK,
	}, nil
}
