package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
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
