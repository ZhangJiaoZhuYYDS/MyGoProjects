package main

import (
	"GormDemo/initDB"
	"fmt"
	"gorm.io/gorm"
)

// User 拥有并属于多种 language，`user_languages` 是连接表
type User4 struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user4_languages;"`
}

type Language struct {
	gorm.Model
	Name string
	Users []User4 `gorm:"many2many:user4_languages;"`
}


func main() {
	initDB.DB.AutoMigrate(&User4{},&Language{})
	// 一个用户使用多种语言
	//l1 := Language{
	//	Name: "中文",
	//}
	//l2 := Language{
	//	Name: "英文",
	//}
	//u1 := User4{
	//	Languages: []Language{l1, l2},
	//}
	//initDB.DB.Create(&u1)

	// 一种语言被多个用户使用
	//u := User4{}
	//l := Language{
	//	Name:  "外星语",
	//	//也可以直接指定创建好的记录的主键
	//	Users: []User4{u, User4{Model: gorm.Model{ID: 1}}},
	//}
	//initDB.DB.Create(&l)
	// 查找
	u4 := User4{}
	initDB.DB.Where("id = ?", 1).Find(&u4)
	fmt.Println(u4)
	initDB.DB.Model(&User4{}).Preload("Languages").Find(&u4)
	fmt.Println(u4)

	u := User4{}
	initDB.DB.Where("id = ?", 2).Find(&u)
	var l []Language
	initDB.DB.Model(&u).Association("Languages").Find(&l)
	fmt.Println(l)




}
