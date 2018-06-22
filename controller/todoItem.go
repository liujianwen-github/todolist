package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"todoList/dao/models"
)

func AddTodo(c *gin.Context) {
	title := c.Request.FormValue("title")
	content := c.Request.FormValue("content")
	status, err := models.AddTodo(&models.TODO{title, content})
	println(status)
	if err != nil {
		//msg := make(map[string]interface{})
		//msg["status"] = 200
		//d, _ := json.Marshal(msg)
		c.String(http.StatusForbidden,"failed")
	} else {
		c.String(http.StatusOK, "add")
	}
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
