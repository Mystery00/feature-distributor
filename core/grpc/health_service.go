package grpc

import (
	"context"
	"feature-distributor/core/pb"
)

type HealthServer struct {
	pb.UnimplementedHealthServiceServer
}

func (h HealthServer) CheckHealth(ctx context.Context, in *pb.CheckHealthRequest) (*pb.CheckHealthResponse, error) {
	return &pb.CheckHealthResponse{
		Success: true,
	}, nil
}
