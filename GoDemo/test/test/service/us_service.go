package service

import (
	"gorm.io/gorm"
	"log"
)

var UsService = usService{}

// usService 业务层
type usService struct {
}

func (t usService) db() *gorm.DB {
	return common.Db
}

// List 分页列表
func (t usService) List(page, size int, v *model.Us) map[string]interface{} {
	// 结果
	var lists []model.Us
	t.db().Model(&v).Where(&v).Order("").Offset((page - 1) * size).Limit(size).Find(&lists)
	// 统计
	var total int64
	t.db().Model(&v).Where(&v).Count(&total)
	data := make(map[string]interface{})
	data["list"] = lists
	data["total"] = total
	return data
}

// One 根据主键Id查询记录
func (t usService) One(id interface{}) model.Us {
	var v model.Us
	db := t.db().Find(&v, id)
	if db.RowsAffected != 1 {
		log.Println("未找到数据！")
	}
	return v
}

// Update 修改记录 true -> 操作成功
func (t usService) Update(v model.Us) bool {
	tx := t.db().Model(&v).Updates(v) //当使用 Model 方法，并且值中有主键值时，主键将会被用于构建条件
	if tx.Error != nil {
		log.Panicln(tx.Error.Error())
		return false
	}
	return true
}

// Insert 插入记录 true -> 操作成功 注: 主键也传递进来的话，那么就会执行更新或插入操作
func (t usService) Insert(v model.Us) bool {
	tx := t.db().Save(&v)
	if tx.Error != nil {
		log.Panicln(tx.Error.Error())
		return false
	}
	return true
}

// Delete 根据主键删除 true -> 操作成功
func (t usService) Delete(ids []int) bool {
	tx := t.db().Delete(model.Us{}, ids)
	if tx.Error != nil {
		log.Panicln(tx.Error.Error())
		return false
	}
	return true
}
