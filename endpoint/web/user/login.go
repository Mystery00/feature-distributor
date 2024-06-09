package user

import (
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"feature-distributor/endpoint/constants"
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginReq struct {
	Username string `json:"username" required:"true" binding:"required"`
	Password string `json:"password" required:"true" binding:"required"`
}

var login gin.HandlerFunc = func(context *gin.Context) {
	var req LoginReq
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}
	client := grpc.GetUserClient()
	response, err := client.CheckLogin(context.Request.Context(), &pb.CheckLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		context.JSON(500, gin.H{"message": err.Error()})
		return
	}
	if response.GetCode() != http.StatusOK {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "username or password is incorrect"})
		return
	}
	token := generateRandomString(req.Username)
	key := fmt.Sprintf("session:%s", token)
	err = redis.Set(context.Request.Context(), key, req.Username, constants.UserSessionExpire)
	if err != nil {
		context.JSON(500, gin.H{"message": err.Error()})
		return
	}
	context.JSON(int(response.GetCode()), gin.H{
		"token":   token,
		"message": "ok",
	})
}

func generateRandomString(username string) string {
	input1 := fmt.Sprintf("%s-%d", username, time.Now().UnixNano())

	sha1Hasher := sha1.New()
	sha1Hasher.Write([]byte(input1))
	sha1Hash := sha1Hasher.Sum(nil)

	input2 := fmt.Sprintf("%d-%s", time.Now().UnixMilli(), sha1Hash)

	sha512Hasher := sha512.New()
	sha512Hasher.Write([]byte(input2))
	sha512Hash := sha512Hasher.Sum(nil)

	randomString := hex.EncodeToString(sha512Hash)
	return randomString
}
