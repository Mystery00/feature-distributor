package group

import (
	"feature-distributor/common/attribute"
	"feature-distributor/common/operation"
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UpdateOptionReq struct {
	GroupId int64    `json:"groupId" binding:"required" validate:"required"`
	Options []option `json:"options" binding:"required" validate:"required"`
}

var updateOption gin.HandlerFunc = func(c *gin.Context) {
	var req UpdateOptionReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info("invalid params", err)
		resp.FailTrans(c, 400, "common.invalid.params")
		return
	}
	//Options检测
	if len(req.Options) == 0 {
		resp.Fail(c, 400, "invalid options")
		return
	}
	//Type检测
	for _, o := range req.Options {
		attrType := attribute.ParseType(o.AttrType)
		if attrType == nil {
			resp.Fail(c, 400, "invalid attr type")
			return
		}
		operationType := operation.ParseType(o.OperationType)
		if operationType == nil {
			resp.Fail(c, 400, "invalid operation type")
			return
		}
		if !operationType.ForAttributeType(*attrType) {
			resp.Fail(c, 400, "invalid operation type for attribute type")
			return
		}
	}
	//保存数据
	client := grpc.GetReqGroupClient()
	options := make([]*pb.ReqGroupOption, 0, len(req.Options))
	for _, o := range req.Options {
		options = append(options, &pb.ReqGroupOption{
			Index:         o.Index,
			AttrType:      o.AttrType,
			AttrName:      o.AttrName,
			OperationType: o.OperationType,
			AttrValue:     o.AttrValue,
		})
	}
	reqGroup, err := client.UpdateReqGroupOption(c, &pb.UpdateReqGroupOptionRequest{
		GroupId: req.GroupId,
		Options: options,
	})
	if err != nil {
		return
	}
	if err != nil {
		grpc.HandleGRPCError(c, err)
		return
	}
	resp.Data(c, gin.H{
		"reqGroupId": reqGroup.GetGroupId(),
	})
}
