package toggle

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"github.com/gin-gonic/gin"
)

type page struct {
	ProjectId int64 `form:"projectId" required:"true" binding:"required"`
	Index     int   `form:"index" required:"true" binding:"required"`
	Size      int   `form:"size" required:"true" binding:"required"`
}

var list gin.HandlerFunc = func(c *gin.Context) {
	var p page
	err := c.ShouldBindQuery(&p)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	client := grpc.GetToggleClient()
	response, err := client.ListToggle(c.Request.Context(), &pb.ListToggleRequest{
		ProjectId: p.ProjectId,
		Index:     int64(p.Index),
		Size:      int64(p.Size),
	})
	if err != nil {
		grpc.HandleGRPCError(c, err)
		return
	}
	resultList := make([]gin.H, 0, len(response.GetToggles()))
	for _, toggle := range response.GetToggles() {
		resultList = append(resultList, gin.H{
			"id":          toggle.Id,
			"title":       toggle.GetTitle(),
			"key":         toggle.GetKey(),
			"enabled":     toggle.GetEnabled(),
			"valueType":   toggle.GetValueType(),
			"description": toggle.GetDescription(),
		})
	}
	c.JSON(200, gin.H{
		"index": p.Index,
		"size":  p.Size,
		"total": response.GetTotal(),
		"list":  resultList,
	})
}
