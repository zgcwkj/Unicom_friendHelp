package routers

import (
	"net/http"

	"github.com/zgcwkj/friendHelp/controllers"
)

// init 初始化路由
func init() {
	http.HandleFunc("/", controllers.Index)                //默认页面
	http.HandleFunc("/Api", controllers.Api)               //接口
	http.HandleFunc("/ApiOneData", controllers.ApiOneData) //一条自动接口
	http.HandleFunc("/Help", controllers.Help)             //帮助页面
	http.HandleFunc("/Test", controllers.Test)             //测试
}
