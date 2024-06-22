package toggle

import (
	"feature-distributor/common/value"
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SaveReq struct {
	ToggleId    *int64     `json:"toggleId"`
	ProjectId   int64      `json:"projectId" binding:"required" validate:"required"`
	Key         string     `json:"key" binding:"required" validate:"required"`
	Enabled     bool       `json:"enabled"`
	Title       string     `json:"title" binding:"required" validate:"required"`
	Description string     `json:"description"`
	ValueType   string     `json:"valueType" binding:"required" validate:"required"`
	Values      []ValueReq `json:"values" binding:"required" validate:"required"`
}

type ValueReq struct {
	Title         string `json:"title" binding:"required" validate:"required"`
	Value         string `json:"value" binding:"required" validate:"required"`
	Description   string `json:"description"`
	Default       bool   `json:"default"`
	DisabledValue bool   `json:"disabledValue"`
}

var save gin.HandlerFunc = func(c *gin.Context) {
	var req SaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info("invalid params", err)
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}
	//ValueType检测
	parseType := value.ParseType(req.ValueType)
	if parseType == nil {
		c.JSON(400, gin.H{"error": "invalid value type"})
		return
	}
	//Values检测
	if len(req.Values) == 0 {
		c.JSON(400, gin.H{"error": "invalid values"})
		return
	}
	defaultValueIndex := -1
	disabledValueIndex := -1
	for i, v := range req.Values {
		if v.Default {
			if defaultValueIndex != -1 {
				c.JSON(400, gin.H{"error": "duplicate default value"})
				return
			}
			defaultValueIndex = i
		}
		if v.DisabledValue {
			if disabledValueIndex != -1 {
				c.JSON(400, gin.H{"error": "duplicate disabled value"})
				return
			}
			disabledValueIndex = i
		}
	}
	if defaultValueIndex == -1 {
		c.JSON(400, gin.H{"error": "default value not found"})
		return
	}
	if disabledValueIndex == -1 {
		c.JSON(400, gin.H{"error": "disabled value not found"})
		return
	}
	//保存数据
	client := grpc.GetToggleClient()
	values := make([]*pb.SaveToggleValue, 0, len(req.Values))
	for _, v := range req.Values {
		values = append(values, &pb.SaveToggleValue{
			Title:       v.Title,
			Value:       v.Value,
			Description: v.Description,
		})
	}
	toggle, err := client.SaveToggle(c.Request.Context(), &pb.SaveToggleRequest{
		ToggleId:      req.ToggleId,
		ProjectId:     req.ProjectId,
		Enabled:       req.Enabled,
		Title:         req.Title,
		Key:           req.Key,
		Description:   req.Description,
		ValueType:     parseType.String(),
		DefaultValue:  int64(defaultValueIndex),
		DisabledValue: int64(disabledValueIndex),
		Values:        values,
	})
	if err != nil {
		if err != nil {
			grpc.HandleGRPCError(c, err)
			return
		}
	}
	c.JSON(200, gin.H{
		"toggleId": toggle.GetId(),
		"title":    toggle.GetTitle(),
		"key":      toggle.GetKey(),
	})
}
