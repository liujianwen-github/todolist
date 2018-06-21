package models

import (
	"github.com/melonws/goweb/libs/logHelper"
	"log"
	"todoList/dao/mongo"
)

type TODO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func AddTodo(t *TODO) (id int64, err error) {
	session, connection := mongo.CreateModel("todoList")
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}
	defer session.Close()
	data := &TODO{
		Title:   t.Title,
		Content: t.Content,
	}
	error := connection.Insert(data)
	if error != nil {
		logHelper.WriteLog("[插入数据失败]", "mongo/error")
		return 0, error
	}
	logHelper.WriteLog("[插入数据失败]", "mongo/notify")
	return 1, nil
}
