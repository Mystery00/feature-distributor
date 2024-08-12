package toggle

import (
	"feature-distributor/common/value"
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CreateReq struct {
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

var create gin.HandlerFunc = func(c *gin.Context) {
	var req CreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info("invalid params", err)
		resp.FailTrans(c, 400, "common.invalid.params")
		return
	}
	//ValueType检测
	parseType := value.ParseType(req.ValueType)
	if parseType == nil {
		resp.FailTrans(c, 400, "toggle.invalid.value.type")
		return
	}
	//Values检测
	if len(req.Values) == 0 {
		resp.FailTrans(c, 400, "toggle.invalid.value")
		return
	}
	defaultValueIndex := 0
	disabledValueIndex := 0
	for i, v := range req.Values {
		if v.Default {
			if defaultValueIndex != 0 {
				resp.FailTrans(c, 400, "toggle.duplicate.default.value")
				return
			}
			defaultValueIndex = i + 1
		}
		if v.DisabledValue {
			if disabledValueIndex != 0 {
				resp.FailTrans(c, 400, "toggle.duplicate.disabled.value")
				return
			}
			disabledValueIndex = i + 1
		}
	}
	if disabledValueIndex == 0 {
		resp.FailTrans(c, 400, "toggle.disabled.value.not.found")
		return
	}
	//保存数据
	client := grpc.GetToggleClient()
	values := make([]*pb.ToggleValue, 0, len(req.Values))
	for _, v := range req.Values {
		values = append(values, &pb.ToggleValue{
			Title:       v.Title,
			Value:       v.Value,
			Description: v.Description,
		})
	}
	toggle, err := client.CreateToggle(c, &pb.CreateToggleRequest{
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
		grpc.HandleGRPCError(c, err)
		return
	}
	resp.Data(c, gin.H{
		"toggleId": toggle.GetId(),
	})
}
