package model

import "gorm.io/gorm"

// 投票信息
type Votes struct {
	gorm.Model
	Title string  `gorm:"column:title;"`
	VotesOpt []*Votesopt `gorm:"-"`
	TotalNumber int `gorm:"column:total_number;default:0"`
	Type string `gorm:"column:type;default:兴趣爱好"`
}

// 投票配置项目
type Votesopt struct {
	gorm.Model
	UserId uint64 `json:"user_id" gorm:"user_id"`
	Sex string	`json:"sex" gorm:"sex"`
	Hobbies string `json:"hobbies" gorm:"column:hobbies"`
	Status int 	`gorm:"column:status;default:0"`
	Type string `gorm:"column:type;default:兴趣爱好"`

}
type VotesVotesopt struct {
	gorm.Model
	VotesId uint64 `gorm:"column:votes_id"`
	VotesoptId uint64 `gorm:"column:votesopt_id"`
	Votesopt *Votesopt `gorm:"-"`

}
// 中间表

