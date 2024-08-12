package project

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var get gin.HandlerFunc = func(c *gin.Context) {
	var p projectId
	err := c.ShouldBindQuery(&p)
	if err != nil {
		logrus.Info("invalid params", err)
		resp.FailTrans(c, 400, "common.invalid.params")
		return
	}
	client := grpc.GetCoreClient()
	project, err := client.GetProject(c, &pb.ProjectRequest{
		Id: p.Id,
	})
	if err != nil {
		if err != nil {
			grpc.HandleGRPCError(c, err)
			return
		}
	}
	c.JSON(200, gin.H{
		"id":         project.GetId(),
		"name":       project.GetName(),
		"key":        project.GetKey(),
		"server_key": project.GetServerKey(),
		"client_key": project.GetClientKey(),
	})
}
