package models

import (
	"../mongo"
	"github.com/melonws/goweb/libs/logHelper"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

type TODO struct {
	Title    string        `bson:"title"`
	Content  string        `bson:"content"`
	Id_      bson.ObjectId `bson:"_id"`
	CreateAt int64         `bson:"create_at"`
}

//添加一条记录
func AddTodo(t *TODO) (status int64, err error) {
	println(t.Title)
	session, connection := mongo.CreateModel("todoList")
	if session == nil {
		return 0, errors.New("没连上啊")
	}
	defer session.Close()
	data := &TODO{
		Id_:      t.Id_,
		Title:    t.Title,
		Content:  t.Content,
		CreateAt: t.CreateAt,
	}
	error := connection.Insert(data)
	if error != nil {
		logHelper.WriteLog("[插入数据失败]", "mongo/error")
		return 0, error
	}
	logHelper.WriteLog("[插入数据成功]", "mongo/notify")
	return 1, nil
}

//获取一条记录
func GetItem(id string) (status int, data interface{}, err error) {
	session, connection := mongo.CreateModel("todoList")
	if session == nil {
		return 0, nil, errors.New("没连上啊")
	}
	defer session.Close()

	println(id)
	//return

	query := bson.M{"_id": bson.ObjectIdHex(id)}
	var values = new(TODO)
	//filter := bson.M{"title": 1, "content": 1}
	error := connection.Find(query).One(&values)
	//fmt.Println(values)
	//println(values)

	if error != nil {
		return 0, [0]int{}, error
	} else {
		return 1, values, nil
	}
}

//更新一条记录
func UpdateItem(id string, title string, content string) (status int, data interface{}, err error) {
	session, connection := mongo.CreateModel("todoList")
	if session == nil {
		return 0, nil, errors.New("没连上啊")
	}
	defer session.Close()
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	update := bson.M{"$set": bson.M{"title": title, "content": content}}
	error := connection.Update(query, update)
	if error != nil {
		return 0, nil, err
	} else {
		return 1, nil, nil
	}
}

//删除一条记录
func DeleteItem(id string) (status int, data interface{}, err error) {
	session, connection := mongo.CreateModel("todoList")
	if session == nil {
		return 0, nil, errors.New("没连上啊")
	}
	defer session.Close()

	query := bson.M{"_id": bson.ObjectIdHex(id)}
	error := connection.Remove(query)
	if error != nil {
		return 0, nil, err
	} else {
		return 1, nil, nil
	}
}

//获取所有记录
type Result struct {
	title   string `bson:"title"`
	content string `bson:"content"`
}

func GetList() (status int, data interface{}, err error) {
	session, connection := mongo.CreateModel("todoList")
	if session == nil {
		return 0, nil, errors.New("没连上啊")
	}
	defer session.Close()
	query := []bson.M{}
	// 接受数据结果（不能用Result直接定义结果，感觉可以，但是不能实现，有待理解
	var values []TODO
	//筛选字段,_id主键默认返回，如不需要返回id，需要显式定义_id为0
	filter := bson.M{"title": 1, "content": 1}
	error := connection.Find(query).Select(filter).All(&values)
	if error != nil {
		logHelper.WriteLog("获取数据失败", "mongo/error")
		return 0, nil, errors.New("获取数据失败")
	}
	return 1, values, nil
}
