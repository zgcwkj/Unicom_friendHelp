package redis

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/pelletier/go-toml"
)

var rdLink string

// reDis Redis的链接配置
type reDis struct {
	server string //地址
	port   string //端口
}

// init 初始化时，把Redis链接上，供后面使用
func init() {
	//读取配置文件
	config, _ := toml.LoadFile("./config/redis.toml")
	//直接读取
	// user := config.Get("redis.server").(string)
	//转换对象后读取
	redisConfig := config.Get("redis").(*toml.Tree)
	reDis := reDis{
		server: redisConfig.Get("server").(string), //地址
		port:   redisConfig.Get("port").(string),   //端口
	}

	rdlink := fmt.Sprintf("%s:%s", reDis.server, reDis.port)
	rdLink = rdlink
}

// SetCode 设置数据
func SetCode(code string, ok bool) {
	conn, err := redis.Dial("tcp", rdLink)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = conn.Do("SET", code, ok)
	defer conn.Close()
	if err != nil {
		log.Fatal("redis set failed:", err)
	}
}

// SetCode 获取数据
func GetCode(code string) bool {
	conn, err := redis.Dial("tcp", rdLink)
	if err != nil {
		log.Fatal(err)
		return false
	}
	ok, err := redis.Bool(conn.Do("GET", code))
	defer conn.Close()
	if err != nil {
		SetCode(code, true)
		return true
	} else {
		return ok
	}
}
