package group

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UpdateReq struct {
	GroupId     int64  `json:"groupId" binding:"required" validate:"required"`
	Title       string `json:"title" binding:"required" validate:"required"`
	Description string `json:"description"`
}

var update gin.HandlerFunc = func(c *gin.Context) {
	var req UpdateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info("invalid params", err)
		resp.FailTrans(c, 400, "common.invalid.params")
		return
	}
	//保存数据
	client := grpc.GetReqGroupClient()
	reqGroup, err := client.UpdateReqGroup(c, &pb.UpdateReqGroupRequest{
		GroupId:     req.GroupId,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		grpc.HandleGRPCError(c, err)
		return
	}
	resp.Data(c, gin.H{
		"reqGroupId": reqGroup.GetGroupId(),
	})
}
