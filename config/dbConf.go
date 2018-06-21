package config

import (
	"gopkg.in/mgo.v2"
	"time"
)

//type DataBase struct {
//	host string
//	port int
//	user string
//	password string
//	db string
//}
var MongoConf = &mgo.DialInfo{
	Addrs:     []string{"47.104.7.232"},
	Direct:    false,
	Timeout:   time.Second * 1,
	Database:  "todo",
	Source:    "todoList",
	Username:  "root",
	Password:  "root",
	PoolLimit: 4096, // Session.SetPoolLimit
}

