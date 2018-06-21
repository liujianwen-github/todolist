package main

import (
	. "./router"
	"todoList/dao/mongo"
)

func main() {
	println("1")
	router := InitRouter()

	_, collection := mongo.CreateModel("todoList")
	if collection == nil {
		println("error")
	}
	router.Run(":8000")
}
