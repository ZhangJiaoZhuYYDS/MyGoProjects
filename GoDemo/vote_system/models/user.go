// @Author zhangjiaozhu 2023/3/18 13:42:00
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserIdentity string `json:"user_identity" form:"user_identity" gorm:"column:user_identity;"`
	Username string 	`json:"username" form:"username" gorm:"column:username" binding:"required"`
	Password string		`json:"password" form:"password" gorm:"column:password" `
	Sex string   `json:"sex" form:"sex" gorm:"column:sex;"`
	Mail string `json:"mail" form:"mail" gorm:"column:mail"`
}
