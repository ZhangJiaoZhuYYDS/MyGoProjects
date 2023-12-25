// @Author zhangjiaozhu 2023/3/14 9:09:00
package dao

import "TouPiaoSystem/model"


// 创建单个用户投票单
func CreateVote( form model.VoteForm, userId uint64) error {

	var votesopt model.Votesopt
	votesopt.Type = form.Type
	votesopt.Sex = form.Sex
	votesopt.UserId = userId

	// 修改用户投票状态
	votesopt.Status = 1

	// 多选值从切片类转为字符串类型
	s :=""
	for _ , v := range form.Hobbies {
		s+= " "
		s+= v
	}
	votesopt.Hobbies = s

	return model.DB.Self.Debug().Create(&votesopt).Error

}

// 获取所有用户投票信息
func GetVoteAll()  []model.Votesopt{
	var users []model.Votesopt
	model.DB.Self.Debug().Model(&model.Votesopt{}).Find(&users)
	return users
}

//获取当前用户投票信息

func GetVoteById(id uint64) []model.Votesopt {
	var user []model.Votesopt
	model.DB.Self.Model(&model.Votesopt{}).Where("user_id = ?",id).Find(&user)
	return user
}
