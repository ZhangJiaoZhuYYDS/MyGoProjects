package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	type UserInfo struct {
		ID uint
		Phone int
		Password string
		NickName string
	}

	db , err := gorm.Open("mysql","root:1234@(127.0.0.1:3306)/db2")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	db.AutoMigrate(&UserInfo{})
	u1 :=UserInfo{
		ID:       1,
		Phone:    123456,
		Password: "123",
		NickName: "456",
	}
	db.Create(&u1)

}
