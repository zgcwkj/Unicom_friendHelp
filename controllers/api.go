package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/zgcwkj/friendHelp/models"
	"github.com/zgcwkj/friendHelp/redis"
	"github.com/zgcwkj/friendHelp/utils"
)

// ApiOneData 单条数据接口
func ApiOneData(resp http.ResponseWriter, req *http.Request) {
	encryptMobile := req.FormValue("encryptMobile")
	if encryptMobile != "" {
		invitationCode := req.FormValue("invitationCode")
		if invitationCode != "" {
			//# Redis直接过滤
			if redis.GetCode(invitationCode) {
				//# Redis直接过滤
				body := utils.ChinaUnicom_GetFriendHelp(encryptMobile, invitationCode)
				strBody := string(body)
				//# 记录到数据库中
				if strings.Contains(strBody, "成功") {
					models.SetCodeFail(invitationCode, true)
				} else {
					models.SetCodeFail(invitationCode, false)
				}
				//# 记录到数据库中
				if len(strBody) > 200 {
					resp.Write([]byte("参与的小伙伴太多"))
				} else {
					log.Println(strBody)
					resp.Write(body)
				}
			} else {
				resp.Write([]byte("这条数据是无效的数据"))
			}
		} else {
			resp.Write([]byte("请求无效"))
		}
	} else {
		resp.Write([]byte("请求无效"))
	}
}

// ApiSetCodeData 提交邀请码接口
func ApiSetCodeData(resp http.ResponseWriter, req *http.Request) {
	invitationCode := req.FormValue("invitationCode")
	if invitationCode != "" {
		count := models.SetCodeData(invitationCode)
		if count > 0 {
			resp.Write([]byte("添加成功"))
		} else if count == -1 {
			resp.Write([]byte("已经存在"))
		} else {
			resp.Write([]byte("添加失败"))
		}
	} else {
		resp.Write([]byte("请求无效"))
	}
}

// ApiGetCodeData 获取邀请码接口
func ApiGetCodeData(resp http.ResponseWriter, req *http.Request) {
	friendDatas := models.GetCodeData() //查询数据库的数据
	b, _ := json.Marshal(friendDatas)   //将数据传换成json格式
	resp.Write(b)                       //输出出去
}
