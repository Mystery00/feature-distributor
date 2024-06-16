package project

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"github.com/gin-gonic/gin"
)

type projectId struct {
	Id int64 `form:"id" required:"true" binding:"required"`
}

var get gin.HandlerFunc = func(c *gin.Context) {
	var p projectId
	err := c.ShouldBindQuery(&p)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	client := grpc.GetCoreClient()
	response, err := client.GetProject(c.Request.Context(), &pb.GetProjectRequest{
		Id: p.Id,
	})
	if err != nil {
		if err != nil {
			grpc.HandleGRPCError(c, err)
			return
		}
	}
	project := response.GetProject()
	c.JSON(200, gin.H{
		"id":         project.GetId(),
		"name":       project.GetName(),
		"key":        project.GetKey(),
		"server_key": project.GetServerKey(),
		"client_key": project.GetClientKey(),
	})
}
