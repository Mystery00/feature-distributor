package grpc

import (
	"feature-distributor/common/alert"
	"net/http"
)

type ErrorResp struct {
	Status    int
	MessageId string
}

var errorCodeMap = map[alert.Code]ErrorResp{
	alert.ServerInternalError: {Status: http.StatusInternalServerError, MessageId: "common.server.error"},
	alert.InvalidParams:       {Status: http.StatusBadRequest, MessageId: "common.invalid.params"},

	alert.ProjectNotExist: {Status: http.StatusBadRequest, MessageId: "project.not.exist"},
	alert.ProjectExist:    {Status: http.StatusBadRequest, MessageId: "project.exist"},

	alert.ToggleNotExist:     {Status: http.StatusBadRequest, MessageId: "toggle.not.exist"},
	alert.ToggleExist:        {Status: http.StatusBadRequest, MessageId: "toggle.exist"},
	alert.InvalidToggleType:  {Status: http.StatusBadRequest, MessageId: "toggle.invalid.value.type"},
	alert.InvalidToggleValue: {Status: http.StatusBadRequest, MessageId: "toggle.invalid.value"},

	alert.ReqGroupNotExist:     {Status: http.StatusBadRequest, MessageId: "req-group.not.exist"},
	alert.ReqGroupExist:        {Status: http.StatusBadRequest, MessageId: "req-group.exist"},
	alert.InvalidOperationType: {Status: http.StatusBadRequest, MessageId: "req-group.invalid.operation.type"},
}

func ReturnErrorMessage(code alert.Code) ErrorResp {
	s, exist := errorCodeMap[code]
	if !exist {
		return ErrorResp{Status: http.StatusInternalServerError, MessageId: "common.unknown.error"}
	}
	return s
}
