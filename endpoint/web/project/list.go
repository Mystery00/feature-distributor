package project

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"github.com/gin-gonic/gin"
)

type page struct {
	Index int `form:"index" required:"true" binding:"required"`
	Size  int `form:"size" required:"true" binding:"required"`
}

var list gin.HandlerFunc = func(c *gin.Context) {
	var p page
	err := c.ShouldBindQuery(&p)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	client := grpc.GetCoreClient()
	response, err := client.ListProjects(c.Request.Context(), &pb.PageRequest{
		Index: int64(p.Index),
		Size:  int64(p.Size),
	})
	if err != nil {
		grpc.HandleGRPCError(c, err)
		return
	}
	resultList := make([]gin.H, 0, len(response.GetProjects()))
	for _, project := range response.GetProjects() {
		resultList = append(resultList, gin.H{
			"id":         project.GetId(),
			"name":       project.GetName(),
			"key":        project.GetKey(),
			"server_key": project.GetServerKey(),
			"client_key": project.GetClientKey(),
		})
	}
	c.JSON(200, gin.H{
		"index": p.Index,
		"size":  p.Size,
		"total": response.GetTotal(),
		"list":  resultList,
	})
}
