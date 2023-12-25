// @Author zhangjiaozhu 2023/3/11 21:55:00
package models

import (
	"Study/config/mysql"
	"gorm.io/gorm"
)

type ProblemBasic struct {
	gorm.Model
	Identity string 	`gorm:"column:identity;type:varchar(36);" json:"identity"`  // 问题表的统一标识
	//CategoryId string 	`gorm:"column:category_id;type:varchar(255);" json:"category_id"` // 分类id,已逗号分割
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id;"` // 关联问题分类表
	Title string `gorm:"column:title;type:varchar(255);" json:"title"` // 文章标题
	Content string `gorm:"column:content;type:text;" json:"content"` // 文章正文
	MaxRuntime int  `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"` //最大运行时间
	MaxMem int `gorm:"column:max_mem;type:int(11);" json:"max_mem"` // 最大内存

}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}
func GetProblemList( keyword , categoryIdentity string)  *gorm.DB{
	 tx := mysql.DB.Debug().Model(ProblemBasic{}).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic")
	 tx.Where("title like ? or content like ?" , "%"+keyword+"%","%"+keyword+"%")
	 if categoryIdentity != "" {
		 tx.Debug().Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			 Where("pc.category_id = (select cb.id from category_basic cb where cb.identity = ?",categoryIdentity)
	 }
	return tx
}