package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Connect(host string, port string, user string, pass string, database string, options string) (db *xorm.Engine, err error) {
	return xorm.NewEngine("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&"+options)
}
