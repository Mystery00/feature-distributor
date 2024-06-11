package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

var corsMiddleware gin.HandlerFunc = cors.New(cors.Config{
	AllowAllOrigins:  true,
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	ExposeHeaders:    []string{},
	AllowCredentials: true,
	MaxAge:           2 * time.Hour,
})
