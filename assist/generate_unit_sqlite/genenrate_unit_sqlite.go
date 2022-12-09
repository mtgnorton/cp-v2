package main

import (
	"gf-admin/utility"

	"github.com/gogf/gf/v2/database/gdb"
)

func main() {

	mysqlDb, err := gdb.New(gdb.ConfigNode{
		Host: "127.0.0.1",
		Port: "3306",
		User: "root",
		Pass: "secret",
		Name: "gf-admin",
		Type: "mysql",
	})

	if err != nil {
		panic(err)
	}
	sqliteDb, err := gdb.New(gdb.ConfigNode{
		Link:    "./gf-admin-unit-base.db",
		Type:    "sqlite",
		Charset: "utf8",
	})

	if err != nil {
		panic(err)
	}

	err = utility.Mysql2sqlite(mysqlDb, sqliteDb)
	if err != nil {
		panic(err)

	}
}
