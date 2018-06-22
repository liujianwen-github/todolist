package main

import (
	. "./router"
)

func main() {
	println("1")
	router := InitRouter()

	//session, _ := mongo.CreateModel("todoList")
	////_, err := models.AddTodo(&models.TODO{"title", "content"})
	//if session == nil {
	//	println("error" )
	//}
	//println("success")
	router.Run(":8000")
}
