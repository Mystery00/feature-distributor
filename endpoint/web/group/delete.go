package group

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var remove gin.HandlerFunc = func(c *gin.Context) {
	var t reqGroupId
	err := c.ShouldBindQuery(&t)
	if err != nil {
		logrus.Info("invalid params", err)
		resp.FailTrans(c, 400, "common.invalid.params")
		return
	}
	client := grpc.GetReqGroupClient()
	_, err = client.DeleteReqGroup(c, &pb.GetReqGroupRequest{
		GroupId: t.GroupId,
	})
	if err != nil {
		if err != nil {
			grpc.HandleGRPCError(c, err)
			return
		}
	}
	resp.Empty(c)
}
