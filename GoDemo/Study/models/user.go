// @Author zhangjiaozhu 2023/3/11 22:06:00
package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity string 	`gorm:"column:identity;type:varchar(36);" json:"identity"`  // 用户表的统一标识
	UserName string `gorm:"column:username;type:varchar(100);" json:"username"`
	Password string `gorm:"column:password;type:varchar(32);" json:"password"`
	Phone string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	Mail string `gorm:"column:mail;type:varchar(100);" json:"mail"`
	FinishProblemNum int64 `gorm:"column:finish_problem_num;type:int(11);" json:"finish_problem_num"` //完成问题的个数
    SubmitNum int64 `gorm:"column:submit_num;type:int(11);" json:"submit_num"`   //提交次数
}

func (table *UserBasic)TableName() string  {
	return "user_basic"
}
