package project

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"github.com/gin-gonic/gin"
)

var remove gin.HandlerFunc = func(c *gin.Context) {
	var p projectId
	err := c.ShouldBindQuery(&p)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	client := grpc.GetCoreClient()
	_, err = client.DeleteProject(c.Request.Context(), &pb.ProjectRequest{
		Id: p.Id,
	})
	if err != nil {
		if err != nil {
			grpc.HandleGRPCError(c, err)
			return
		}
	}
	c.Status(204)
}
