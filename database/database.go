package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var db *sql.DB

func InitDB() {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	name := viper.GetString("database.name")

	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", user, password, host, port, name) // 某些 MySQL 驱动参数（如 parseTime=true）必须加上，否则 DATETIME/TIMESTAMP 也会被当作字符串处理，Scan 时就会报错。

	var err error
	db, err = sql.Open("mysql", source)
	if err != nil {
		panic("数据库打开失败" + err.Error())
	}

	if err := db.Ping(); err != nil {
		panic("数据库连接失败" + err.Error())
	}
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Println("数据库连接关闭出错:", err)
		} else {
			log.Println("数据库连接关闭成功")
		}
	}
}
