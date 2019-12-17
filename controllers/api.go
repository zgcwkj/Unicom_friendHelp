package controllers

import (
	"net/http"
	"strings"

	"github.com/zgcwkj/friendHelp/models"
	"github.com/zgcwkj/friendHelp/utils"
)

// Api 接口
func Api(resp http.ResponseWriter, req *http.Request) {
	encryptMobile := req.FormValue("encryptMobile")
	if encryptMobile != "" {
		invitationCode := req.FormValue("invitationCode")
		if invitationCode != "" {
			invitationCodes := strings.Split(invitationCode, "<br/>")
			for _, v := range invitationCodes {
				utils.ChinaUnicom_GetFriendHelp(encryptMobile, v)
			}
			resp.Write([]byte("正在处理"))
		} else {
			resp.Write([]byte("请求无效"))
		}
	} else {
		resp.Write([]byte("请求无效"))
	}
}

// ApiOneData 单条数据接口
func ApiOneData(resp http.ResponseWriter, req *http.Request) {
	encryptMobile := req.FormValue("encryptMobile")
	if encryptMobile != "" {
		invitationCode := req.FormValue("invitationCode")
		if invitationCode != "" {
			body := utils.ChinaUnicom_GetFriendHelp(encryptMobile, invitationCode)
			resp.Write(body)
		} else {
			resp.Write([]byte("请求无效"))
		}
	} else {
		resp.Write([]byte("请求无效"))
	}
}

func Test(resp http.ResponseWriter, req *http.Request) {
	models.GetRow()
}
