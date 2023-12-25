// @Author zhangjiaozhu 2023/3/12 12:18:00
package models

import (
	"gorm.io/gorm"
)

type ProblemCategory struct {
	gorm.Model
	ProblemId     uint           `gorm:"column:problem_id;" json:"problem_id"`           // 问题的ID
	CategoryId    uint           `gorm:"column:category_id;" json:"category_id"`         // 分类的ID
	// category_basic表的字段id关联problem_category的字段category_id  他们的字段类型都是uint，但是迁移表显示无法添加外键约束
	CategoryBasic *CategoryBasic `gorm:"foreignKey:category_id;references:id" json:"category_basic"`  // 关联分类的基础信息表
}

func (*ProblemCategory)TableName() string  {
	return "problem_category"
}