package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //引用 MySql 包
	"github.com/pelletier/go-toml"
)

// Db 数据库操作内存
var Db *sql.DB

// mySQL 数据的链接配置
type mySQL struct {
	server   string //数据库地址
	port     string //数据库端口
	database string //数据库名称
	name     string //数据库用户
	pass     string //数据库密码
}

// init 初始化时，把数据链接上，供后面使用
func init() {
	//读取配置文件
	config, _ := toml.LoadFile("./config/mysql.toml")
	//直接读取
	// user := config.Get("mysql.server").(string)
	//转换对象后读取
	mysqlConfig := config.Get("mysql").(*toml.Tree)
	mysql := mySQL{
		server:   mysqlConfig.Get("server").(string),   //数据库地址
		port:     mysqlConfig.Get("port").(string),     //数据库端口
		database: mysqlConfig.Get("database").(string), //数据库名称
		name:     mysqlConfig.Get("name").(string),     //数据库用户
		pass:     mysqlConfig.Get("pass").(string),     //数据库密码
	}

	dblink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", mysql.name, mysql.pass, mysql.server, mysql.port, mysql.database)
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		log.Fatal(err)
		return
	}
	Db = db
	// defer db.Close()
}
