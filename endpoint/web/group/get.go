package group

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sort"
)

var get gin.HandlerFunc = func(c *gin.Context) {
	var t reqGroupId
	err := c.ShouldBindQuery(&t)
	if err != nil {
		logrus.Info("invalid params", err)
		resp.FailTrans(c, 400, "common.invalid.params")
		return
	}
	client := grpc.GetReqGroupClient()
	reqGroup, err := client.GetReqGroup(c, &pb.GetReqGroupRequest{
		GroupId: t.GroupId,
	})
	if err != nil {
		grpc.HandleGRPCError(c, err)
		return
	}
	indexs := make([]int64, 0)
	for _, groupOption := range reqGroup.GetOptions() {
		if contains(indexs, groupOption.GetIndex()) {
			continue
		}
		indexs = append(indexs, groupOption.GetIndex())
	}
	sort.Slice(indexs, func(i, j int) bool {
		return indexs[i] < indexs[j]
	})

	optionMap := make(map[int64][]option)
	for _, opt := range reqGroup.GetOptions() {
		if l, exist := optionMap[opt.GetIndex()]; exist {
			optionMap[opt.GetIndex()] = append(l, option{
				Index:         opt.GetIndex(),
				AttrType:      opt.GetAttrType(),
				AttrName:      opt.GetAttrName(),
				OperationType: opt.GetOperationType(),
				AttrValue:     opt.GetAttrValue(),
			})
			continue
		}
		optionMap[opt.GetIndex()] = []option{{
			Index:         opt.GetIndex(),
			AttrType:      opt.GetAttrType(),
			AttrName:      opt.GetAttrName(),
			OperationType: opt.GetOperationType(),
			AttrValue:     opt.GetAttrValue(),
		}}
	}
	options := make([][]option, 0)
	for _, index := range indexs {
		options = append(options, optionMap[index])
	}
	resp.Data(c, gin.H{
		"title":       reqGroup.GetTitle(),
		"key":         reqGroup.GetKey(),
		"description": reqGroup.GetDescription(),
		"options":     options,
	})
}

func contains[T int | int64](slice []T, item T) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}
