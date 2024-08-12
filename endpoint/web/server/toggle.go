package server

import (
	"feature-distributor/common/value"
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"github.com/gin-gonic/gin"
	"strings"
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
	valueType := strings.ToLower(t.ToggleType)
	client := grpc.GetToggleClient()
	response, err := client.GetToggleValue(c, &pb.GetToggleValueRequest{
		ReqUser:    t.ReqUser.buildReqUser(),
		ProjectKey: projectKey,
		ToggleKey:  t.ToggleKey,
		ToggleType: valueType,
	})
	if err != nil {
		grpc.HandleGRPCError(c, err)
		return
	}
	detectValue, err := value.AutoDetectValue(response.GetResultValue(), valueType)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"value": detectValue})
}
