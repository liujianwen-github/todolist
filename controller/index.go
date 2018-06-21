package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"github.com/melonws/goweb/libs/logHelper"
)

func IndexApi(c *gin.Context) {
	logHelper.WriteLog("[通知请求][请求数据][1,23,4,5]","notify/access")

	c.String(http.StatusOK,"It works")
}