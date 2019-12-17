package controllers

import (
	"io/ioutil"
	"net/http"
	"os"
)

// Index 默认页面
func Index(resp http.ResponseWriter, req *http.Request) {
	f, _ := os.Open("./views/index.html") //读取文件
	buf, _ := ioutil.ReadAll(f)           //将文件转换成数据
	resp.Write(buf)                       // 输出到页面
}

// Help 帮助页面
func Help(resp http.ResponseWriter, req *http.Request) {
	f, _ := os.Open("./views/help.html") //读取文件
	buf, _ := ioutil.ReadAll(f)          //将文件转换成数据
	resp.Write(buf)                      // 输出到页面
}
