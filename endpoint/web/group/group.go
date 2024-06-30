package group

import "github.com/gin-gonic/gin"

func Handle(engine *gin.RouterGroup) {
	engine.POST("/v1", create)
	//engine.GET("/list/v1", list)
	//engine.GET("/v1", get)
	//engine.DELETE("/v1", remove)
}

type reqGroupId struct {
	GroupId int64 `form:"groupId" required:"true" binding:"required"`
}

type option struct {
	Index         int64  `json:"index" binding:"required" validate:"required"`
	AttrType      string `json:"attrType" binding:"required" validate:"required"`
	AttrName      string `json:"attrName" binding:"required" validate:"required"`
	OperationType string `json:"operationType" binding:"required" validate:"required"`
	AttrValue     string `json:"attrValue" binding:"required" validate:"required"`
}
