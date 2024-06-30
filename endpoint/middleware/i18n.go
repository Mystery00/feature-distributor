package middleware

import (
	"feature-distributor/endpoint/i18n"
	"github.com/gin-gonic/gin"
)

var i18nMiddleware gin.HandlerFunc = func(c *gin.Context) {
	accept := c.GetHeader("Accept-Language")
	localizer := i18n.NewLocalizer(accept)
	c.Set(i18n.GinLocalizerKey, localizer)
}
