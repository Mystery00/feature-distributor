package project

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
)

var remove gin.HandlerFunc = func(c *gin.Context) {
	var p projectId
	err := c.ShouldBindQuery(&p)
	if err != nil {
		resp.Err(c, 400, err)
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
	resp.Empty(c)
}
