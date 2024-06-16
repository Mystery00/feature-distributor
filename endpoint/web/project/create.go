package project

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CreateReq struct {
	Name string `json:"name" required:"true" binding:"required"`
	Key  string `json:"key" required:"true" binding:"required"`
}

var create gin.HandlerFunc = func(c *gin.Context) {
	var req CreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info("invalid params", err)
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}
	client := grpc.GetCoreClient()
	response, err := client.SaveProject(c.Request.Context(), &pb.SaveProjectRequest{
		Name: req.Name,
		Key:  req.Key,
	})
	if err != nil {
		if err != nil {
			grpc.HandleGRPCError(c, err)
			return
		}
	}
	c.JSON(200, gin.H{
		"id":   response.GetProject().GetId(),
		"name": response.GetProject().GetName(),
		"key":  response.GetProject().GetKey(),
	})
}
