// @Author zhangjiaozhu 2023/3/7 19:24:00
package models

import "time"

type Post struct {
	PostID uint64 `json:"post_id" gorm:"post_id"`
	AuthorID uint64 `json:"author_id"`
	CommunityID uint64 `json:"community_id" binding:"required"`
	Status int32 `json:"status"`
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	CreateTime time.Time `json:"-"`

}
