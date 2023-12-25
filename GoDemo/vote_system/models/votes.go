// @Author zhangjiaozhu 2023/3/19 16:25:00
package models

import (
	"gorm.io/gorm"
)

// 投票信息
type Votes struct {
	gorm.Model
	Title string  `gorm:"column:title;"`
	VotesOpt []Votesopt `gorm:"foreignKey:Type;references:Type"`
	TotalNumber int `gorm:"column:total_number;default:0"`
	Type string `gorm:"column:type;"`
	During int `gorm:"column:during;default:1"`
}

// 投票配置项目
type Votesopt struct {
	gorm.Model
	UserIdentity string `json:"user_identity" form:"user_identity" gorm:"column:user_identity;"`
	Result int	`json:"result" gorm:"result"`  // 是否赞成 0 赞成  1 反对
	Hobbies string `json:"hobbies" gorm:"column:hobbies"`  // 兴趣爱好
	Status int 	`gorm:"column:status;default:0"`   // 投票状态 0 未投票 1 已投票
	Type string `gorm:"column:type;"` // 投票类型

}
type VotesVotesopt struct {
	gorm.Model
	VotesId uint64 `gorm:"column:votes_id"`   // 票id
	VotesoptId uint64 `gorm:"column:votesopt_id"` // 投票者id
	Votesopt *Votesopt `gorm:"-"`

}