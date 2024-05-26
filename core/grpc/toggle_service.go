package grpc

import (
	"context"
	"encoding/json"
	"feature-distributor/core/pb"
	"github.com/open-feature/go-sdk/openfeature"
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
