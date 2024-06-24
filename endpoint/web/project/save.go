package project

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SaveReq struct {
	Id   *int64 `json:"id"`
	Name string `json:"name" required:"true" binding:"required"`
	Key  string `json:"key" required:"true" binding:"required"`
}

var save gin.HandlerFunc = func(c *gin.Context) {
	var req SaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info("invalid params", err)
		resp.Fail(c, 400, "invalid params")
		return
	}
	client := grpc.GetCoreClient()
	project, err := client.SaveProject(c.Request.Context(), &pb.SaveProjectRequest{
		ProjectId: req.Id,
		Name:      req.Name,
		Key:       req.Key,
	})
	if err != nil {
		if err != nil {
			grpc.HandleGRPCError(c, err)
			return
		}
	}
	resp.Data(c, gin.H{
		"id":   project.GetId(),
		"name": project.GetName(),
		"key":  project.GetKey(),
	})
}
