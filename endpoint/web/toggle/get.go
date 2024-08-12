package toggle

import (
	"feature-distributor/common/value"
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var get gin.HandlerFunc = func(c *gin.Context) {
	var t toggleId
	err := c.ShouldBindQuery(&t)
	if err != nil {
		logrus.Info("invalid params", err)
		resp.FailTrans(c, 400, "common.invalid.params")
		return
	}
	client := grpc.GetToggleClient()
	toggle, err := client.GetToggle(c, &pb.GetToggleRequest{
		Id: t.Id,
	})
	if err != nil {
		if err != nil {
			grpc.HandleGRPCError(c, err)
			return
		}
	}
	values := make([]gin.H, 0, len(toggle.Values))
	for _, v := range toggle.Values {
		values = append(values, gin.H{
			"title":         v.GetTitle(),
			"value":         v.GetValue(),
			"description":   v.GetDescription(),
			"default":       v.GetId() == toggle.DefaultValue,
			"disabledValue": v.GetId() == toggle.DisabledValue,
		})
	}
	c.JSON(200, gin.H{
		"id":          toggle.GetId(),
		"projectId":   toggle.GetProjectId(),
		"enabled":     toggle.GetEnabled(),
		"title":       toggle.GetTitle(),
		"key":         toggle.GetKey(),
		"description": toggle.GetDescription(),
		"valueType":   *value.ParseType(toggle.ValueType),
		"values":      values,
	})
}
