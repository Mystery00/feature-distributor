package client

import (
	"context"
	"feature-distributor/common/value"
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
	value, err := getToggle(c, t.ReqUser, projectKey, t.ToggleKey, t.ToggleType)
	if err != nil {
		grpc.HandleGRPCError(c, err)
		return
	}
	c.JSON(200, gin.H{"value": value})
}

func getToggle(ctx context.Context, r reqUser, projectKey, toggleKey, toggleType string) (any, error) {
	client := grpc.GetToggleClient()
	response, err := client.GetToggleValue(ctx, &pb.GetToggleValueRequest{
		ReqUser:    r.buildReqUser(),
		ProjectKey: projectKey,
		ToggleKey:  toggleKey,
		ToggleType: toggleType,
	})
	if err != nil {
		return nil, err
	}
	return value.AutoDetectValue(toggleType, response.GetResultValue())
}
