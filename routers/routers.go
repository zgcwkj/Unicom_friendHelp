package routers

import (
	"net/http"

	"github.com/zgcwkj/Unicom_friendHelp/controllers"
)

// init 初始化路由
func init() {
	http.HandleFunc("/", controllers.Index)                        //默认页面
	http.HandleFunc("/ApiOneData", controllers.ApiOneData)         //一条自动接口
	http.HandleFunc("/ApiSetCodeData", controllers.ApiSetCodeData) //提交邀请码接口
	http.HandleFunc("/ApiGetCodeData", controllers.ApiGetCodeData) //获取邀请码接口
	http.HandleFunc("/Help", controllers.Help)                     //帮助页面
}
