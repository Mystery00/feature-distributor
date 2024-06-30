package resp

import (
	"feature-distributor/endpoint/i18n"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func FailTrans(c *gin.Context, code int, id string) {
	Fail(c, code, i18n.Translate(c, id))
}

func Fail(c *gin.Context, code int, message string) {
	c.JSON(code, Result{
		Code: code,
		Msg:  message,
		Data: nil,
	})
	c.Abort()
}

func Err(c *gin.Context, code int, err error) {
	logrus.Errorf("error: %v", err)
	c.JSON(code, Result{
		Code: code,
		Msg:  err.Error(),
		Data: nil,
	})
	c.Abort()
}

func Empty(c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		Code: http.StatusOK,
		Msg:  "success",
		Data: nil,
	})
}

func Ok(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Result{
		Code: http.StatusOK,
		Msg:  message,
		Data: nil,
	})
}

func OkData(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, Result{
		Code: http.StatusOK,
		Msg:  message,
		Data: data,
	})
}

func Data(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Result{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	})
}
