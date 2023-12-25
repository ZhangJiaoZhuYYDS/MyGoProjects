// @Author zhangjiaozhu 2023/3/7 17:24:00
package mysql

import (
	"Buble/models"
	"database/sql"
	"fmt"
)

func GetCommunityList() (communityList []*models.Community , err error) {
	sqlStr := `select community_id , community_name from community`
	err = _db.Raw(sqlStr).Scan(&communityList).Debug().Error
	if err != nil {
		fmt.Println("查询为空")
		err = nil
	}
	return
}

func GetCommunityNameByID(idStr string) (community *models.Community, err error) {
	community = new(models.Community)
	sqlStr := `select community_id, community_name
	from community
	where community_id = ?`
	err = _db.Raw(sqlStr,idStr).Scan(&community).Debug().Error
	if err == sql.ErrNoRows {
		return
	}
	if err != nil {
		return
	}
	return
}



func GetCommunityById( id uint64) (*models.Community,error) {
	var community models.Community
	sqlStr := `select community_id , community_name , introduction ,create_time from community where community_id = ?`
	err := _db.Raw(sqlStr,id).Scan(&community).Debug().Error
	return &community , err
}