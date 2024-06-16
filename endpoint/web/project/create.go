package project

import (
	"feature-distributor/common/alert"
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CreateReq struct {
	Name string `json:"name" required:"true" binding:"required"`
	Key  string `json:"key" required:"true" binding:"required"`
}

var create gin.HandlerFunc = func(context *gin.Context) {
	var req CreateReq
	err := context.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info("invalid params", err)
		context.JSON(400, gin.H{"error": "invalid params"})
		return
	}
	client := grpc.GetCoreClient()
	response, err := client.SaveProject(context.Request.Context(), &pb.SaveProjectRequest{
		Name: req.Name,
		Key:  req.Key,
	})
	if err != nil {
		if c := alert.Convert(err); c != nil {
			if *c == alert.ProjectExist {
				context.JSON(400, gin.H{"error": "项目唯一标识符已存在"})
				return
			}
		}
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{
		"id":   response.GetProject().GetId(),
		"name": response.GetProject().GetName(),
		"key":  response.GetProject().GetKey(),
	})
}
