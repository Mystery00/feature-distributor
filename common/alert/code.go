package alert

type Code int

const (
	InvalidParams Code = 1000

	ProjectNotExist Code = 1001
	ProjectExist    Code = 1002

	ToggleNotExist     Code = 1003
	ToggleExist        Code = 1004
	InvalidToggleType  Code = 1005
	InvalidToggleValue Code = 1006

	ReqGroupExist        Code = 1007
	ReqGroupNotExist     Code = 1008
	InvalidOperationType Code = 1009

	ServerInternalError Code = 5000
)
