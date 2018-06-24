package controller

import (
	"../dao/models"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)
//新增一条todo
func AddTodo(c *gin.Context) {
	title := c.Request.FormValue("title")
	content := c.Request.FormValue("content")
	_, err := models.AddTodo(&models.TODO{title, content, bson.NewObjectId(), time.Now().Unix()})
	if err != nil {
		//msg := make(map[string]interface{})
		//msg["status"] = 200
		//d, _ := json.Marshal(msg)
		c.String(http.StatusForbidden, "failed")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"stataus": 200,
			"message": "success",
			"data":    nil,
		})
	}
}
//删除一条todo
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	status, data, err := models.DeleteItem(id)
	checkData(c, status, data, err)
	//c.String(http.StatusOK, "delete")
}
//更新一条todo
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	fmt.Println(title, content)
	status, data, err := models.UpdateItem(id, title, content)
	checkData(c, status, data, err)
	//c.String(http.StatusOK, "update")
}
//获取一条todo
func GetTodo(c *gin.Context) {
	// Query 获取参数体参数；Param 获取rest参数；DefaultQuery("page":"0")  赋默认值
	id := c.Param("id")
	println("获取参数:", id)
	println("查询参数：", bson.ObjectIdHex(id))
	fmt.Println(bson.ObjectIdHex(id))
	status, data, err := models.GetItem(id)
	//c.String(http.StatusOK,"return")
	//return
	println(status)
	checkData(c, status, data, err)
}
//获取全部todo
func GetList(c *gin.Context) {
	status, data, err := models.GetList()
	checkData(c, status, data, err)

}

// 检查操作结果是否正确并返回相应的response
func checkData(c *gin.Context, status int, data interface{}, err error) {
	println(c, status, data, err)
	if err != nil {
		println(err.Error())
		c.String(http.StatusBadGateway, err.Error())
	} else {
		if status != 1 {
			c.JSON(http.StatusBadGateway, gin.H{
				"status":  500,
				"message": "failed",
				"data":    [0]int{},
			})
		} else {
			// 返回数据方式 1、c.string，将json转为string；2、c.json返回map；3、c.json返回gin.H对象，应该是一个map转为json的方法
			c.JSON(http.StatusOK, gin.H{
				"status":  200,
				"message": "success",
				"data":    data,
			})
		}
	}
}
