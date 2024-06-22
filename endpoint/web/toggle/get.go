package toggle

import (
	"feature-distributor/common/value"
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"github.com/gin-gonic/gin"
)

var get gin.HandlerFunc = func(c *gin.Context) {
	var t toggleId
	err := c.ShouldBindQuery(&t)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	client := grpc.GetToggleClient()
	toggle, err := client.GetToggle(c.Request.Context(), &pb.GetToggleRequest{
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
			"default":       v.Id == toggle.DefaultValue,
			"disabledValue": v.Id == toggle.DisabledValue,
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
