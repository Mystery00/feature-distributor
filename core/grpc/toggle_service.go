package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"feature-distributor/common/alert"
	"feature-distributor/core/db/enum"
	"feature-distributor/core/db/query"
	"feature-distributor/core/pb"
	"github.com/open-feature/go-sdk/openfeature"
	"gorm.io/gorm"
)

var clientMap = make(map[string]*openfeature.Client)

func getClient(projectKey string) *openfeature.Client {
	if client, ok := clientMap[projectKey]; ok {
		return client
	}
	client := openfeature.NewClient(projectKey)
	clientMap[projectKey] = client
	return client
}

type ToggleServer struct {
	pb.UnimplementedToggleServiceServer
}

func (s *ToggleServer) ListToggle(ctx context.Context, in *pb.ListToggleRequest) (*pb.ListToggleResponse, error) {
	t := query.Toggle
	tc := t.WithContext(ctx)

	offset := int((in.GetIndex() - 1) * in.GetSize())
	limit := int(in.GetSize())
	list, total, err := tc.Where(t.ProjectID.Eq(in.GetProjectId())).FindByPage(offset, limit)
	if err != nil {
		return nil, err
	}
	toggleList := make([]*pb.ListItemToggle, 0)
	for _, item := range list {
		toggleList = append(toggleList, &pb.ListItemToggle{
			Id:          item.ID,
			Enabled:     item.Enable,
			Title:       item.Title,
			Key:         item.Key,
			Description: item.Description,
			ValueType:   enum.ValueTypeEnum(item.ValueType).String(),
		})
	}
	return &pb.ListToggleResponse{
		Total:   total,
		Toggles: toggleList,
	}, nil
}

func (s *ToggleServer) GetToggle(ctx context.Context, in *pb.GetToggleRequest) (*pb.Toggle, error) {
	t := query.Toggle
	tv := query.ToggleValue
	p := query.Project
	tc := t.WithContext(ctx)
	tvc := tv.WithContext(ctx)
	pc := p.WithContext(ctx)

	toggle, err := tc.Where(t.ID.Eq(in.GetId())).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, alert.Error(alert.ToggleNotExist)
		}
		return nil, err
	}
	values, err := tvc.Where(tv.ToggleID.Eq(toggle.ID)).Find()
	if err != nil {
		return nil, err
	}
	valueList := make([]*pb.ToggleValue, 0)
	for _, item := range values {
		valueList = append(valueList, &pb.ToggleValue{
			Id:          item.ID,
			ToggleId:    item.ToggleID,
			Title:       item.Title,
			Value:       item.Value,
			Description: item.Description,
			CreateTime:  item.CreateTime.UnixMilli(),
		})
	}
	project, err := pc.Where(p.ID.Eq(toggle.ProjectID)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, alert.Error(alert.ProjectNotExist)
		}
		return nil, err
	}
	return &pb.Toggle{
		Id:            toggle.ID,
		ProjectId:     toggle.ProjectID,
		ProjectKey:    project.Key,
		Enabled:       toggle.Enable,
		Title:         toggle.Title,
		Key:           toggle.Key,
		Description:   toggle.Description,
		ValueType:     enum.ValueTypeEnum(toggle.ValueType).String(),
		DefaultValue:  toggle.DefaultValue,
		DisabledValue: toggle.ReturnValueWhenDisable,
		CreateTime:    toggle.CreateTime.UnixMilli(),
		UpdateTime:    toggle.UpdateTime.UnixMilli(),
		Values:        valueList,
	}, nil
}

func (s *ToggleServer) GetBoolToggle(ctx context.Context, in *pb.BoolToggleRequest) (*pb.BoolToggleResponse, error) {
	client := getClient(in.GetProjectKey())
	attributes := make(map[string]interface{})
	for k, v := range in.GetReqUser().GetAttributes() {
		attributes[k] = v
	}
	c := openfeature.NewEvaluationContext(in.GetReqUser().GetRolloutKey(), attributes)
	value, err := client.BooleanValue(ctx, in.GetToggleKey(), in.GetDefaultValue(), c)
	if err != nil {
		return nil, err
	}
	return &pb.BoolToggleResponse{
		ToggleKey:   in.GetToggleKey(),
		ResultValue: value,
	}, nil
}

func (s *ToggleServer) GetStringToggle(ctx context.Context, in *pb.StringToggleRequest) (*pb.StringToggleResponse, error) {
	client := getClient(in.GetProjectKey())
	attributes := make(map[string]interface{})
	for k, v := range in.GetReqUser().GetAttributes() {
		attributes[k] = v
	}
	c := openfeature.NewEvaluationContext(in.GetReqUser().GetRolloutKey(), attributes)
	value, err := client.StringValue(ctx, in.GetToggleKey(), in.GetDefaultValue(), c)
	if err != nil {
		return nil, err
	}
	return &pb.StringToggleResponse{
		ToggleKey:   in.GetToggleKey(),
		ResultValue: value,
	}, nil
}

func (s *ToggleServer) GetFloatToggle(ctx context.Context, in *pb.FloatToggleRequest) (*pb.FloatToggleResponse, error) {
	client := getClient(in.GetProjectKey())
	attributes := make(map[string]interface{})
	for k, v := range in.GetReqUser().GetAttributes() {
		attributes[k] = v
	}
	c := openfeature.NewEvaluationContext(in.GetReqUser().GetRolloutKey(), attributes)
	value, err := client.FloatValue(ctx, in.GetToggleKey(), float64(in.GetDefaultValue()), c)
	if err != nil {
		return nil, err
	}
	return &pb.FloatToggleResponse{
		ToggleKey:   in.GetToggleKey(),
		ResultValue: float32(value),
	}, nil
}

func (s *ToggleServer) GetIntToggle(ctx context.Context, in *pb.IntToggleRequest) (*pb.IntToggleResponse, error) {
	client := getClient(in.GetProjectKey())
	attributes := make(map[string]interface{})
	for k, v := range in.GetReqUser().GetAttributes() {
		attributes[k] = v
	}
	c := openfeature.NewEvaluationContext(in.GetReqUser().GetRolloutKey(), attributes)
	value, err := client.IntValue(ctx, in.GetToggleKey(), in.GetDefaultValue(), c)
	if err != nil {
		return nil, err
	}
	return &pb.IntToggleResponse{
		ToggleKey:   in.GetToggleKey(),
		ResultValue: value,
	}, nil
}

func (s *ToggleServer) GetJsonToggle(ctx context.Context, in *pb.JsonToggleRequest) (*pb.JsonToggleResponse, error) {
	client := getClient(in.GetProjectKey())
	attributes := make(map[string]interface{})
	for k, v := range in.GetReqUser().GetAttributes() {
		attributes[k] = v
	}
	c := openfeature.NewEvaluationContext(in.GetReqUser().GetRolloutKey(), attributes)
	value, err := client.ObjectValue(ctx, in.GetToggleKey(), in.GetDefaultValue(), c)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return &pb.JsonToggleResponse{
		ToggleKey:   in.GetToggleKey(),
		ResultValue: string(bytes),
	}, nil
}
