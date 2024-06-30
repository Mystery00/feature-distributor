package group

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type page struct {
	Index int `form:"index" required:"true" binding:"required"`
	Size  int `form:"size" required:"true" binding:"required"`
}

var list gin.HandlerFunc = func(c *gin.Context) {
	var p page
	err := c.ShouldBindQuery(&p)
	if err != nil {
		logrus.Info("invalid params", err)
		resp.FailTrans(c, 400, "common.invalid.params")
		return
	}
	client := grpc.GetReqGroupClient()
	response, err := client.ListReqGroup(c.Request.Context(), &pb.ListReqGroupRequest{
		Index: int64(p.Index),
		Size:  int64(p.Size),
	})
	if err != nil {
		grpc.HandleGRPCError(c, err)
		return
	}
	resultList := make([]gin.H, 0, len(response.GetGroups()))
	for _, group := range response.GetGroups() {
		resultList = append(resultList, gin.H{
			"groupId":     group.GetGroupId(),
			"title":       group.GetTitle(),
			"key":         group.GetKey(),
			"description": group.GetDescription(),
		})
	}
	resp.Data(c, gin.H{
		"total": response.GetTotal(),
		"list":  resultList,
	})
}
