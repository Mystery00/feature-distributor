package grpc

import (
	"feature-distributor/common/alert"
	"net/http"
)

type ErrorResp struct {
	Status int
	Msg    string
}

var errorCodeMap = map[alert.Code]ErrorResp{
	alert.InvalidParams:     {Status: http.StatusBadRequest, Msg: "参数错误"},
	alert.ProjectNotExist:   {Status: http.StatusBadRequest, Msg: "项目不存在"},
	alert.ProjectExist:      {Status: http.StatusBadRequest, Msg: "项目已存在"},
	alert.ToggleNotExist:    {Status: http.StatusBadRequest, Msg: "配置项不存在"},
	alert.ToggleExist:       {Status: http.StatusBadRequest, Msg: "配置项已存在"},
	alert.InvalidToggleType: {Status: http.StatusBadRequest, Msg: "配置项类型错误"},

	alert.ServerInternalError: {Status: http.StatusInternalServerError, Msg: "服务器内部错误"},
}

func ReturnErrorMessage(code alert.Code) ErrorResp {
	s, exist := errorCodeMap[code]
	if !exist {
		return ErrorResp{Status: http.StatusInternalServerError, Msg: "未知错误"}
	}
	return s
}
