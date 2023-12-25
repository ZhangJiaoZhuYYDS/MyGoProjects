// @Author zhangjiaozhu 2023/3/9 12:54:00
package main

import (
	"GormDemo/initDB"
	"gorm.io/gorm"
	"log"
)

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name      string
	CompanyID int // 默认情况下， CompanyID 被隐含地用来在 User 和 Company 之间创建一个外键关系
	Company   Company `gorm:"foreignKey:CompanyID"`   // 自定义外键
	// Company   Company `gorm:"references:Name"` // 使用 Name  重写引用
}

type Company struct {
	ID   int
	Name string
}

func main() {
	//user 里面有 company表的结构 所以只需要自动迁移user表即可
	initDB.DB.AutoMigrate(&User{})
	// 创建
	//c := Company{
	//	ID:   1,
	//	Name: "dachang",
	//}

	//u := User{
	//	Name:    "linzy",
	//	Company: c,
	//}
	// initDB.DB.Create(&u)
	// 查询
	var u1 User
	initDB.DB.Model(&User{}).First(&u1)
	log.Println(u1)

	initDB.DB.Where("id = ?",1).Take(&u1)
	log.Println(u1)
		// 预加载
		initDB.DB.Model(&User{}).Preload("Company").First(&u1)
		log.Println(u1)
		// 查找关联
		var c1 Company
		err := initDB.DB.Model(&u1).Association("Company").Find(&c1)
		if err != nil {
			return
		}
		log.Println(c1)

}

