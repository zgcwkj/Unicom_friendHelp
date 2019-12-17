package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
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

// GetRow 获取一行数据 friendData
// 随机获取 chickensoup 表的一行数据
func GetRow() {
	// rows, err := Db.Query("SELECT * FROM frienddata ORDER BY RAND() LIMIT 1")
	rows, err := Db.Query("SELECT * FROM frienddata ORDER BY RAND() LIMIT 10")
	if err != nil {
		log.Fatalln(rows, err)
	}
	log.Println(rows)
	defer rows.Close()
	cloumns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	values := make([]sql.RawBytes, len(cloumns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(cloumns[i], ": ", value)
		}
		fmt.Println("------------------")
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	// return chickensoup
}
