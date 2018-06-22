package models

import (
	"todoList/dao/mongo"
	"github.com/pkg/errors"
	"github.com/melonws/goweb/libs/logHelper"
)

type TODO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func AddTodo(t *TODO) (status int64, err error) {
	println(t.Title)
	session, connection := mongo.CreateModel("todoList")
	if session == nil {
		return 0, errors.New("没连上啊")
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
	logHelper.WriteLog("[插入数据成功]", "mongo/notify")
	return 1, nil
}
