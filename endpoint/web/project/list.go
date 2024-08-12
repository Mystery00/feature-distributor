package project

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
)

var list gin.HandlerFunc = func(c *gin.Context) {
	client := grpc.GetCoreClient()
	response, err := client.ListProjects(c, &pb.PageRequest{
		Index: 1,
		Size:  100,
	})
	if err != nil {
		grpc.HandleGRPCError(c, err)
		return
	}
	resultList := make([]gin.H, 0, len(response.GetProjects()))
	for _, project := range response.GetProjects() {
		resultList = append(resultList, gin.H{
			"id":        project.GetId(),
			"name":      project.GetName(),
			"key":       project.GetKey(),
			"serverKey": project.GetServerKey(),
			"clientKey": project.GetClientKey(),
		})
	}
	resp.Data(c, resultList)
}
