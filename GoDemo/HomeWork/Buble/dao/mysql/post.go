// @Author zhangjiaozhu 2023/3/7 20:14:00
package mysql

import (
	"Buble/models"
	"github.com/lexkong/log"
)

func CreatePost(post *models.Post) (err error)  {
	sqlStr := `insert into post(post_id,title,content,author_id,community_id) values (?,?,?,?,?)`
	err = _db.Raw(sqlStr, post.PostID, post.Title, post.Content, post.AuthorID, post.CommunityID).Debug().Error
	if err != nil {
		log.Info("创建文章失败")
		return
	}
	return
}

func GetPostById()  {
	
}