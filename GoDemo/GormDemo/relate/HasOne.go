package main

import (
	"GormDemo/initDB"
	"gorm.io/gorm"
	"log"
)

// User 有一张 CreditCard，UserID 是外键
type User2 struct {
	gorm.Model
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	User2ID uint
}


func main() {
	// 迁移两个数据库
	initDB.DB.AutoMigrate(&User2{},&CreditCard{})
	// 创建
	//c := CreditCard{
	//	Number:  "123456",
	//}
	//u := User2{
	//	CreditCard: c ,
	//}
	//initDB.DB.Create(&u)
	// 查询
	var u2 User2
	initDB.DB.Model(&User2{}).First(&u2)
	log.Println(u2)
	// 预加载
	initDB.DB.Model(&User2{}).Preload("CreditCard").First(&u2)
	log.Println(u2)
	var cd CreditCard
	initDB.DB.Debug().Model(&u2).Association("CreditCard").Find(&cd)
	log.Println(cd)

}
