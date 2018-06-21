package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func AddTodo(c *gin.Context) {
	c.String(http.StatusOK, "add")
}
func DeleteTodo(c *gin.Context) {
	c.String(http.StatusOK, "delete")
}
func UpdateTodo(c *gin.Context) {
	c.String(http.StatusOK, "update")
}
func GetTodo(c *gin.Context) {
	c.String(http.StatusOK, "get")
}
func GetList(c *gin.Context)  {
	c.String(http.StatusOK, "getList")
}
