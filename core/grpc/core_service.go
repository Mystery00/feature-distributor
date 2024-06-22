package grpc

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"feature-distributor/common/alert"
	"feature-distributor/core/db/model"
	"feature-distributor/core/db/query"
	"feature-distributor/core/notify"
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
		resultList = append(resultList, convertToProject(project))
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
		Project: convertToProject(project),
	}, nil
}

func (s *CoreServer) GetProject(ctx context.Context, in *pb.ProjectRequest) (*pb.Project, error) {
	p := query.Project
	pc := p.WithContext(ctx)
	project, err := pc.Where(p.ID.Eq(in.GetId())).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, alert.Error(alert.ProjectNotExist)
		}
		return nil, err
	}
	return convertToProject(project), nil
}

func (s *CoreServer) SaveProject(ctx context.Context, in *pb.SaveProjectRequest) (*pb.Project, error) {
	p := query.Project
	pc := p.WithContext(ctx)
	var project *model.Project
	if in.ProjectId != nil {
		//更新
		_, err := pc.Where(p.ID.Eq(in.GetProjectId())).First()
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, alert.Error(alert.ProjectNotExist)
			}
			return nil, err
		}
		_, err = pc.Where(p.ID.Eq(in.GetProjectId())).Updates(model.Project{
			Name: in.GetName(),
		})
		if err != nil {
			return nil, err
		}
		project, _ = pc.Where(p.ID.Eq(in.GetProjectId())).First()
	} else {
		//查询是否存在相同的key
		pro, err := pc.Where(p.Key.Eq(in.GetKey())).First()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if pro != nil {
			return nil, alert.Error(alert.ProjectExist)
		}
		pro = &model.Project{
			Name:      in.GetName(),
			Key:       in.GetKey(),
			ServerKey: generateKey(in.GetKey(), "toggle"),
			ClientKey: generateKey(in.GetKey(), "client"),
		}
		err = pc.Save(pro)
		if err != nil {
			return nil, err
		}
		project = pro
	}
	notify.ProjectChange(*project)
	return convertToProject(project), nil
}

func (s *CoreServer) DeleteProject(ctx context.Context, in *pb.ProjectRequest) (*pb.Project, error) {
	p := query.Project
	pc := p.WithContext(ctx)
	project, err := pc.Where(p.ID.Eq(in.GetId())).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, alert.Error(alert.ProjectNotExist)
		}
		return nil, err
	}
	_, err = pc.Where(p.ID.Eq(in.GetId())).Delete()
	if err != nil {
		return nil, err
	}
	notify.ProjectChange(*project)
	return convertToProject(project), nil
}

func convertToProject(p *model.Project) *pb.Project {
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
