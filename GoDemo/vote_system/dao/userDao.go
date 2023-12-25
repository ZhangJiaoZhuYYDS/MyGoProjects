// @Author zhangjiaozhu 2023/3/18 16:08:00
package dao

import (
	"vote_system/models"
	"vote_system/utils/md5"
	"vote_system/utils/mysql"
	"vote_system/utils/uuid"
)
// 注册保存用户信息
func Save(user models.User)  (error){
	user.UserIdentity=uuid.GetUUID()
	user.Password = md5.GetMd5(user.Password)
	return  mysql.DB.Create(&user).Error
}
// 登录 根据用户名查询用户信息
func GetUserByUsername(username string) (models.User ,error) {
	var user models.User
	err := mysql.DB.Model(&models.User{}).Where("username = ?", username).First(&user).Error
	return user, err
}

func FindUserStatus(userIdentity string , voteType string) (int,error) {
	var status int
	err := mysql.DB.Model(&models.Votesopt{}).Select("status").Where("user_identity = ? AND type = ?", userIdentity, voteType).First(&status).Error
	if err != nil {
		return -1 ,err
	}
	return status,nil
}