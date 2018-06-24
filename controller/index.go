package controller

import (
	"github.com/melonws/goweb/libs/logHelper"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type Response struct {
	status  int
	message string
	data    interface{}
}

func IndexApi(c *gin.Context) {
	logHelper.WriteLog("[通知请求][请求数据][1,23,4,5]", "notify/access")

	c.String(http.StatusOK, "It works")
}
