// @Author zhangjiaozhu 2023/3/11 22:12:00
package models

import (
	"gorm.io/gorm"
	"time"
)

type CategoryBasic struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Identity string 	`gorm:"column:identity;type:varchar(36);" json:"identity"`  // 分类表的统一标识
	Name string `gorm:"column:name;type:varchar(100);" json:"name"`  // 分类名称

}

func (table *CategoryBasic)TableName() string {
	return "category_basic"
}