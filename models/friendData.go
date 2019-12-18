package models

import (
	"log"
	"time"

	"github.com/zgcwkj/friendHelp/redis"
)

// friendData 映射 friendData 表的结构数据
type friendData struct {
	ID         int       `json:"id"`          //主键
	Code       string    `json:"code"`        //邀请码
	Fail       string    `json:"fail"`        //失败次数
	Creaator   string    `json:"creaator"`    //创建者信息
	IsDelete   bool      `json:"is_delete"`   //删除否
	CreateTime time.Time `json:"create_time"` //创建时间
}

// SetCodeData 提交邀请码
func SetCodeData(code string) int64 {
	selectCode := "" //存储查询的数据（避免有重复的数据）
	Db.QueryRow("SELECT CODE FROM frienddata WHERE code = ?", code).Scan(&selectCode)
	if selectCode == "" {
		res, err := Db.Exec("INSERT INTO frienddata(code) VALUE(?)", code)
		if err != nil {
			log.Fatalln(res, err)
		}
		// defer res.Close()
		// res.LastInsertId() //得到修改数据的主键
		// res.RowsAffected() //得到影响行数
		count, _ := res.RowsAffected()
		return count
	}
	return -1
}

// GetCodeData 获取邀请码
func GetCodeData() []friendData {
	rows, err := Db.Query("SELECT id, code, fail, create_time  FROM frienddata WHERE fail < 5 and create_time >= DATE_SUB(NOW(), INTERVAL 12 HOUR) ORDER BY RAND() LIMIT 10")
	if err != nil {
		log.Fatalln(rows, err)
	}
	defer rows.Close()
	cloumns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	frienddatas := make([]friendData, len(cloumns))
	for rows.Next() {
		frienddata := friendData{}
		// rows.Scan(&frienddata...)
		rows.Scan(&frienddata.ID, &frienddata.Code, &frienddata.Fail, &frienddata.CreateTime)
		frienddatas = append(frienddatas, frienddata)
	}
	return frienddatas
}

// GetCodeData 获取邀请码
func SetCodeFail(code string, ok bool) int64 {
	res, err := Db.Exec("UPDATE frienddata SET fail = fail+1 WHERE code = ?", code) //更改无效的数据
	if err != nil {
		log.Fatalln(res, err)
	}
	count, _ := res.RowsAffected()
	if ok {
		if count == 0 {
			res, err := Db.Exec("INSERT INTO frienddata(code) VALUE(?)", code) //存储有效的数据
			if err != nil {
				log.Fatalln(res, err)
			}
			count, _ := res.RowsAffected()
			return count
		}
	}
	//存储到Redis，直接不处理
	redis.SetCode(code, ok)
	return -1
}
