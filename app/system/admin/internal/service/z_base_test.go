package service

import (
	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/mattn/go-sqlite3"
)

var ctx = gctx.New()

func init() {
	//err := utility.InitUnit()
	//
	//if err != nil {
	//	panic(err)
	//}
}
