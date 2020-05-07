package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// ChinaUnicom_GetFriendHelp 联通年终活动接口方法
func ChinaUnicom_GetFriendHelp(encryptMobile string, invitationCode string) []byte {
	// url := "https://m.client.10010.com/DoubleCard_Pro/static/doubleCard/friendHelp?"
	// https://m.client.10010.com/DoubleCard_Pro/static/doubleCard/activityIndexDu?encryptMobile=c99288d5f025f56709cc6c80c4df095f&invitationCode=
	url := "https://m.client.10010.com/DoubleCard_Pro/static/doubleCard/"
	data := "encryptMobile=" + encryptMobile + "&invitationCode=" + invitationCode
	client := &http.Client{}
	reqest, err := http.NewRequest("POST", url+"actFriendHelp?"+data, nil)
	reqest.Header.Add("Cookie", "")
	reqest.Header.Add("Referer", url+"activityIndexDu?"+data)
	reqest.Header.Add("Origin", "https://m.client.10010.com")
	reqest.Header.Add("Access-Control-Allow-Credentials", "true")
	reqest.Header.Add("Access-Control-Allow-Origin", "https://img.client.10010.com")
	reqest.Header.Add("Accept", "application/json, text/plain, */*")
	reqest.Header.Add("User-Agent", "Mozilla/5.0")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	//处理返回结果
	response, _ := client.Do(reqest)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}
