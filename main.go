package main

import (
	"log"
	"net/http"

	_ "github.com/zgcwkj/Unicom_friendHelp/routers"
)

func main() {
	log.Println("作者：zgcwkj")                  //输出信息
	log.Println("说明：仅可用于联通年终活动")    //输出信息
	log.Println("时间：2019-12-18 22:10")        //输出信息
	log.Println("打开 http://127.0.0.1:845 访问")//输出地址
	http.ListenAndServe(":845", nil)             //监听端口
}
