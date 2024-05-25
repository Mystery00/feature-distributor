package grpc

import (
	"context"
	"feature-distributor/core/db/query"
	"feature-distributor/core/pb"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedCoreServiceServer
}

func (s *Server) CheckProject(ctx context.Context, in *pb.CheckProjectRequest) (*pb.CheckProjectResponse, error) {
	if in.ClientKey == nil && in.ServerKey == nil {
		return nil, errors.New("invalid params")
	}
	p := query.Project
	pc := p.WithContext(ctx)
	if in.ClientKey != nil {
		pc = pc.Where(p.ClientKey.Eq(*in.ClientKey))
	} else if in.ServerKey != nil {
		pc = pc.Where(p.ServerKey.Eq(*in.ServerKey))
	}
	project, err := pc.First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pb.CheckProjectResponse{
				IsProjectExist: false,
			}, nil
		}
		return nil, err
	}
	return &pb.CheckProjectResponse{
		IsProjectExist: true,
		ProjectId:      &project.ID,
		ProjectKey:     &project.Key,
		ProjectName:    &project.Name,
	}, nil
}
