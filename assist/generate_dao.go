package main

import (
	"fmt"
	"os/exec"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

func main() {

	table := "forum_posts" // 需要生成dao层的表

	ctx := gctx.New()

	command := "gf gen dao"

	if table != "" { // 使用的配置文件
		command = `gf gen dao --gf.gcfg.file=config-local.toml  --gf.debug=1  -t ` + table
	}

	cmd := exec.Command("/bin/zsh", "-c", command)
	var output []byte
	var err error
	if output, err = cmd.Output(); err != nil {
		g.Log().Fatalf(ctx, "执行gf gen dao 错误: %s,%s", err, output)
		return
	}
	g.Dump(output)

	rootPath := gfile.Pwd()

	Dirs := map[string]string{
		"/app/service/internal/dao":          "/app/dao/",
		"/app/service/internal/dao/internal": "/app/dao/internal/",
		"/app/service/internal/do":           "/app/dto/",
	}
	for tempSource, tempDst := range Dirs {
		_, err = gfile.ScanDirFileFunc(rootPath+tempSource, "", false, func(sourcePath string) string {

			filename := gfile.Basename(sourcePath)

			if gstr.HasPrefix(filename, "_") {
				filename = gstr.TrimLeft(filename, "_")
			}

			dstPath := rootPath + tempDst + filename

			content := gfile.GetContents(sourcePath)

			content = gstr.Replace(content, "gf-admin/app/service/internal/dao/internal", "gf-admin/app/dao/internal")

			//  PostsColumns 变量名首字母大写问题
			if gstr.Contains(tempSource, "/app/service/internal/dao/internal") {
				// 下划线转为驼峰
				ucFirstFieldName := gstr.CaseCamel(gstr.Replace(filename, ".go", ""))
				lcFirstFieldName := gstr.LcFirst(ucFirstFieldName)
				fmt.Println(ucFirstFieldName, lcFirstFieldName)
				content = gstr.Replace(content, "var "+ucFirstFieldName, "var "+lcFirstFieldName)

				content = gstr.Replace(content, "columns: "+ucFirstFieldName, "columns: "+lcFirstFieldName)
			}

			// app/dto 命令空间由 do 转为dto
			if gstr.Contains(tempSource, "/app/service/internal/do") {
				content = gstr.Replace(content, "package do", "package dto")
			}

			err = gfile.PutContents(dstPath, content)

			if err != nil {
				g.Log().Fatalf(ctx, "写入文件错误：%s", err)
			}

			g.Log().Infof(ctx, "%s 移动到 %s \n", sourcePath, dstPath)

			return sourcePath
		})
		g.Dump(err)
	}

	err = gfile.Remove(rootPath + "/app/service")

	// 去除/model/entity 下生成的文件名以下划线开头
	_, err = gfile.ScanDirFileFunc(rootPath+"/app/model/entity", "", false, func(sourcePath string) string {

		filename := gfile.Basename(sourcePath)

		if gstr.HasPrefix(filename, "_") {
			dstPath := rootPath + "/app/model/entity/" + gstr.Replace(filename, "_", "", 1)
			err = gfile.Move(sourcePath, dstPath)
			if err != nil {
				g.Log().Errorf(ctx, "移动文件错误：%s,sourcePath: %s, dstPath: %s", err, sourcePath, dstPath)
			} else {
				g.Log().Infof(ctx, "移动文件成功,sourcePath: %s, dstPath: %s", sourcePath, dstPath)
			}
		}
		return ""
	})

	g.Dump(err)
}
