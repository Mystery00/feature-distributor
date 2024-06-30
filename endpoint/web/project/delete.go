package project

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var remove gin.HandlerFunc = func(c *gin.Context) {
	var p projectId
	err := c.ShouldBindQuery(&p)
	if err != nil {
		logrus.Info("invalid params", err)
		resp.FailTrans(c, 400, "common.invalid.params")
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
