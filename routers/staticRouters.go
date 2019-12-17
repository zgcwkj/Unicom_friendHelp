package routers

import "net/http"

// init 初始化静态文件
func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

}
