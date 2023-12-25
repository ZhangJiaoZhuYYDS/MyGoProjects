// @Author zhangjiaozhu 2023/3/7 17:11:00
package service

import (
	"Buble/dao/mysql"
	"Buble/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityById(id uint64) ( *models.Community, error)  {
	return mysql.GetCommunityById(id)
}