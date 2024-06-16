package grpc

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"feature-distributor/common/alert"
	"feature-distributor/core/db/model"
	"feature-distributor/core/db/query"
	"feature-distributor/core/pb"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type CoreServer struct {
	pb.UnimplementedCoreServiceServer
}

func (s *CoreServer) ListProjects(ctx context.Context, in *pb.PageRequest) (*pb.ProjectPageResponse, error) {
	p := query.Project
	pc := p.WithContext(ctx)
	offset := int((in.GetIndex() - 1) * in.GetSize())
	limit := int(in.GetSize())
	list, total, err := pc.FindByPage(offset, limit)
	if err != nil {
		return nil, err
	}
	resultList := make([]*pb.Project, 0, len(list))
	for _, project := range list {
		resultList = append(resultList, convertTo(project))
	}
	return &pb.ProjectPageResponse{
		Total:    total,
		Projects: resultList,
	}, nil
}

func (s *CoreServer) CheckProject(ctx context.Context, in *pb.CheckProjectRequest) (*pb.CheckProjectResponse, error) {
	if in.ClientKey == nil && in.ServerKey == nil {
		return nil, alert.Error(alert.InvalidParams)
	}
	p := query.Project
	pc := p.WithContext(ctx)
	if in.ClientKey != nil {
		pc = pc.Where(p.ClientKey.Eq(in.GetClientKey()))
	} else if in.ServerKey != nil {
		pc = pc.Where(p.ServerKey.Eq(in.GetServerKey()))
	}
	project, err := pc.First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pb.CheckProjectResponse{
				Project: nil,
			}, nil
		}
		return nil, err
	}
	return &pb.CheckProjectResponse{
		Project: convertTo(project),
	}, nil
}

func (s *CoreServer) GetProject(ctx context.Context, in *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	p := query.Project
	pc := p.WithContext(ctx)
	project, err := pc.Where(p.ID.Eq(in.GetId())).First()
	if err != nil {
		return nil, err
	}
	return &pb.GetProjectResponse{
		Project: convertTo(project),
	}, nil
}

func (s *CoreServer) SaveProject(ctx context.Context, in *pb.SaveProjectRequest) (*pb.GetProjectResponse, error) {
	p := query.Project
	pc := p.WithContext(ctx)
	//查询是否存在相同的key
	project, err := pc.Where(p.Key.Eq(in.GetKey())).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if project != nil {
		return nil, alert.Error(alert.ProjectExist)
	}
	project = &model.Project{
		Name:      in.GetName(),
		Key:       in.GetKey(),
		ServerKey: generateKey(in.GetKey(), "server"),
		ClientKey: generateKey(in.GetKey(), "client"),
	}
	err = pc.Save(project)
	if err != nil {
		return nil, err
	}
	return &pb.GetProjectResponse{
		Project: convertTo(project),
	}, nil
}

func convertTo(p *model.Project) *pb.Project {
	if p == nil {
		return nil
	}
	return &pb.Project{
		Id:         p.ID,
		Name:       p.Name,
		Key:        p.Key,
		ServerKey:  p.ServerKey,
		ClientKey:  p.ClientKey,
		CreateTime: p.CreateTime.UnixMilli(),
		UpdateTime: p.UpdateTime.UnixMilli(),
	}
}

func generateKey(key, salt string) string {
	input := fmt.Sprintf("%s-%d-%s", key, time.Now().UnixMilli(), salt)
	h := sha256.New()
	h.Write([]byte(input))
	hash := h.Sum(nil)
	randomString := hex.EncodeToString(hash)
	return randomString
}
