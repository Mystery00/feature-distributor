package user

import (
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"feature-distributor/endpoint/constants"
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/redis"
	"feature-distributor/endpoint/web/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type LoginReq struct {
	Username string `json:"username" required:"true" binding:"required"`
	Password string `json:"password" required:"true" binding:"required"`
}

var login gin.HandlerFunc = func(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Info("invalid params", err)
		resp.FailTrans(c, 400, "common.invalid.params")
		return
	}
	client := grpc.GetUserClient()
	response, err := client.CheckLogin(c, &pb.CheckLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		grpc.HandleGRPCError(c, err)
		return
	}
	if response.GetCode() != http.StatusOK {
		resp.FailTrans(c, 400, "user.wrong.password")
		return
	}
	token := generateRandomString(req.Username)
	key := fmt.Sprintf("session:%s", token)
	session := make(map[string]any)
	session["userId"] = response.GetUserId()
	session["username"] = req.Username
	sessionJson, err := json.Marshal(session)
	if err != nil {
		resp.Err(c, 500, err)
		return
	}
	err = redis.Set(c, key, string(sessionJson), constants.UserSessionExpire)
	if err != nil {
		resp.Err(c, 500, err)
		return
	}
	resp.Data(c, token)
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
