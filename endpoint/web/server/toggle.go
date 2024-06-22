package server

import (
	"context"
	"errors"
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"github.com/gin-gonic/gin"
)

type toggleRequest struct {
	ReqUser    reqUser `json:"reqUser" required:"true"`
	ToggleKey  string  `json:"toggleKey" required:"true"`
	ToggleType string  `json:"toggleType" required:"true"`
}

var toggle gin.HandlerFunc = func(c *gin.Context) {
	projectKey := c.GetString("projectKey")
	t := &toggleRequest{}
	err := c.ShouldBindBodyWithJSON(t)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
		})
		c.Abort()
		return
	}
	value, err := getToggle(c.Request.Context(), t.ReqUser, projectKey, t.ToggleKey, t.ToggleType)
	if err != nil {
		grpc.HandleGRPCError(c, err)
		return
	}
	c.JSON(200, gin.H{"value": value})
}

func getToggle(ctx context.Context, r reqUser, projectKey, toggleKey, toggleType string) (any, error) {
	client := grpc.GetToggleClient()
	switch toggleType {
	case "bool":
		response, err := client.GetBoolToggle(ctx, &pb.BoolToggleRequest{
			ReqUser:    r.buildReqUser(),
			ProjectKey: projectKey,
			ToggleKey:  toggleKey,
		})
		if err != nil {
			return nil, err
		}
		return response.GetResultValue(), nil
	case "string":
		response, err := client.GetStringToggle(ctx, &pb.StringToggleRequest{
			ReqUser:    r.buildReqUser(),
			ProjectKey: projectKey,
			ToggleKey:  toggleKey,
		})
		if err != nil {
			return nil, err
		}
		return response.GetResultValue(), nil
	case "float":
		response, err := client.GetFloatToggle(ctx, &pb.FloatToggleRequest{
			ReqUser:    r.buildReqUser(),
			ProjectKey: projectKey,
			ToggleKey:  toggleKey,
		})
		if err != nil {
			return nil, err
		}
		return response.GetResultValue(), nil
	case "int":
		response, err := client.GetIntToggle(ctx, &pb.IntToggleRequest{
			ReqUser:    r.buildReqUser(),
			ProjectKey: projectKey,
			ToggleKey:  toggleKey,
		})
		if err != nil {
			return nil, err
		}
		return response.GetResultValue(), nil
	case "json":
		response, err := client.GetJsonToggle(ctx, &pb.JsonToggleRequest{
			ReqUser:    r.buildReqUser(),
			ProjectKey: projectKey,
			ToggleKey:  toggleKey,
		})
		if err != nil {
			return nil, err
		}
		return response.GetResultValue(), nil
	default:
		return nil, errors.New("invalid toggle type")
	}
}
