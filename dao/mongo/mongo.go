package mongo

import (
	"github.com/melonws/goweb/libs/logHelper"
	"gopkg.in/mgo.v2"
	"time"
)

type TODO struct {
	_id string
}
func connect() *mgo.Session {

	mongoDialInfo := &mgo.DialInfo{
		Addrs:     []string{"47.104.7.232:27017"},
		Direct:    false,
		Timeout:   time.Second * 1,
		Database:  "todo",
		Source:    "todoList",
		Username:  "root",
		Password:  "root",
		PoolLimit: 4096, // Session.SetPoolLimit
	}

	session,err := mgo.DialWithInfo(mongoDialInfo)

	if err != nil {
		logHelper.WriteLog("mongo 链接失败"+err.Error(),"mongodb/error")
		return nil
	}

	return session
}
/**
 * 通过数据库连接后，拿到对应的session(会话) 和 collection(集合)
 * 便于操作
 * 返回后的操作，记得先 defer session.Close()
 */
func CreateModel(collectionName string) (*mgo.Session,*mgo.Collection) {

	session := connect()

	if session == nil {
		return nil,nil
	}

	session.SetMode(mgo.Monotonic,true)

	collection := session.DB("todo").C(collectionName)

	return session,collection
}
