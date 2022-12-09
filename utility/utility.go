package utility

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gogf/gf/v2/os/gcfg"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	"github.com/gogf/gf/v2/database/gdb"
)

func EncryptPassword(username, password string) string {
	return gmd5.MustEncrypt(username + password)
}

// GetServerPath 获取运行时的服务器路径
func GetServerPath() string {
	return gfile.Join(gfile.Pwd(), g.Cfg().MustGet(gctx.New(), "server.serverRoot").String())
}

// MySQL2SQLite converts MySQL database to SQLite database.
func Mysql2sqlite(mysqlDb gdb.DB, sqliteDb gdb.DB) error {

	ctx := gctx.New()

	tables, err := mysqlDb.Tables(ctx)
	if err != nil {
		return err
	}

	for _, table := range tables {
		_, err = sqliteDb.Exec(ctx, fmt.Sprintf("DROP TABLE IF EXISTS `%s`", table))
		if err != nil {
			return err
		}

		// Create table
		_, err = sqliteDb.Exec(ctx, fmt.Sprintf(`
	CREATE TABLE %s (
		id          INTEGER       PRIMARY KEY AUTOINCREMENT
									UNIQUE
									NOT NULL
	);
`, table))
		if err != nil {
			return err
		}

		// add columns
		fields, err := mysqlDb.TableFields(ctx, table)
		if err != nil {
			return err
		}

		for _, field := range fields {

			if field.Name == "id" {
				continue
			}
			fieldType := gstr.Replace(field.Type, "unsigned", "")
			_, err = sqliteDb.Exec(ctx, fmt.Sprintf(`
	ALTER TABLE %s ADD COLUMN %s %s;
`, table, field.Name, fieldType))
			if err != nil {
				return err
			}
		}

		// if table is log , ignore
		if gstr.Contains(table, "_log") {
			continue
		}

		// write data
		rows, err := mysqlDb.GetAll(ctx, "select * from "+table)

		if err != nil {
			return err
		}

		if len(rows) == 0 {
			return nil
		}

		rowsLine, err := sqliteDb.Insert(ctx, table, rows)
		if err != nil {
			return err
		}
		rowsAffected, err := rowsLine.RowsAffected()
		if err != nil {
			return err
		}
		fmt.Printf("table:%s , insert rows:%d \n", table, rowsAffected)

		if err != nil {
			return err
		}

	}

	return nil
}

// GetRootPath 获取项目根路径
//Go没有提供接口让我们区分程序是go run还是go build执行，但我们可以换个思路来实现：
//根据go run的执行原理，我们得知它会源代码编译到系统TEMP或TMP环境变量目录中并启动执行；
//那我们可以直接在程序中对比os.Executable()获取到的路径是否与环境变量TEMP设置的路径相同，
//如果相同，说明是通过go run启动的，因为当前执行路径是在TEMP目录；不同的话自然是go build
//的启动方式。
func GetRootPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	fmt.Println(dir, tmpDir)
	if strings.Contains(dir, tmpDir) {
		dir = getCurrentAbPathByCaller()
	}

	for {
		// 说明是根目录
		if gfile.Exists(gfile.Join(dir, "config")) {
			return dir
		} else { //继续向上查找
			dir = gfile.Dir(dir)
			if dir == "." || dir == "/" {
				return ""
			}
		}

	}
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// InitUnit 设置单元测试初始化操作,使用sqlite3数据库替代mysql
func InitUnit(useSqlite ...bool) error {

	// 部分情况下会出现项目配置文件路径错误的情况，手动添加查询的配置路径
	err := g.Cfg().GetAdapter().(*gcfg.AdapterFile).AddPath(gfile.Join(GetRootPath(), "config"))
	if err != nil {
		return err
	}

	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-unit.toml")

	if len(useSqlite) > 0 && useSqlite[0] == false {
		return nil
	}

	// 初始化sqlite文件
	sqliteBasePath := gfile.Join(GetRootPath(), "gf-admin-unit-base.db")
	sqliteInstancePath := gfile.Join(GetRootPath(), "gf-admin-unit-base-instance.db")
	if !gfile.Exists(sqliteBasePath) {
		return errors.New("gf-admin-unit-base.db not exists")
	}
	if gfile.Exists(sqliteInstancePath) {
		err := gfile.Remove(sqliteInstancePath)
		if err != nil {
			return err
		}
	}

	err = gfile.CopyFile(sqliteBasePath, sqliteInstancePath)
	if err != nil {
		return err
	}
	// 动态设置配置文件使用的数据库
	err = g.Cfg().GetAdapter().(*gcfg.AdapterFile).Set("database.link", "sqlite:"+sqliteInstancePath)

	return err

}
