package grpc

import (
	"context"
	"feature-distributor/core/db/model"
	"feature-distributor/core/notify"
	"feature-distributor/core/pb"
)

type HealthServer struct {
	pb.UnimplementedHealthServiceServer
}

func (h HealthServer) CheckHealth(ctx context.Context, in *pb.CheckHealthRequest) (*pb.CheckHealthResponse, error) {
	notify.ProjectChange(model.Project{ID: 999, Key: "test_test"})
	return &pb.CheckHealthResponse{
		Success: true,
	}, nil
}
