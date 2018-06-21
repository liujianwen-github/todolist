package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"todoList/dao/models"
)

func AddTodo(c *gin.Context) {
	title := c.Request.FormValue("title")
	content := c.Request.FormValue("content")
	_, err := models.AddTodo(&models.TODO{title, content})
	if err != nil {
		c.String(http.StatusForbidden, "failed")
	}
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
func GetList(c *gin.Context) {
	c.String(http.StatusOK, "getList")
}
