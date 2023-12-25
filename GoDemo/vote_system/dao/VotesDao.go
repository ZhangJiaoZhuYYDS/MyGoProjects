// @Author zhangjiaozhu 2023/3/19 18:39:00
package dao

import (
	"vote_system/models"
	"vote_system/utils/mysql"
)

func CreatVote(vote models.Votesopt)  error{
	return mysql.DB.Create(&vote).Error
}

func GetVotesByType(votesType string) (models.Votes, int64,error)  {
	var votes models.Votes
	var count int64
	err := mysql.DB.Preload("VotesOpt").Model(&models.Votes{}).Count(&count).Where("type = ?", votesType).Find(&votes).Error
	return votes,count,err
}

func GetVotesoptByType(s , votesType string) (models.Votesopt,error) {
	var usr  models.Votesopt
	err := mysql.DB.Where("user_identity = ? AND type = ?", s, votesType).First(&usr).Error
	return usr,err
}
func FindVotesByType(s , s2 string) bool {
	var number int64
	mysql.DB.Model(models.Votes{}).Debug().Where("type = ? or title = ?",s,s2).Count(&number)
	if number > 0 {
		return false
	}
	return true
}

func NewVote(vote *models.Votes) bool {
	result := FindVotesByType(vote.Type, vote.Title)
	if result {
		mysql.DB.Create(vote)
		return true
	}
	return false
}