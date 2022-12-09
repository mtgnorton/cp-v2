package main

import (
	"gf-admin/app/system/admin"
	"gf-admin/app/system/index"
	_ "gf-admin/boot"

	"github.com/gogf/gf/v2/os/gctx"
)

//初始化admin用户

func main() {

	go index.Run(gctx.New())
	//
	admin.Run(gctx.New())
}
