// @Author zhangjiaozhu 2023/3/7 19:49:00
package service

import (
	"Buble/commons/snowflake"
	"Buble/dao/mysql"
	"Buble/models"
	"github.com/lexkong/log"
)

func CreatePost(post *models.Post) error {
	// 1 生成post_id
	id := snowflake.GenID()
	uId := uint64(id)
	post.PostID = uId
	// 2 创建帖子，保存到数据库
	err := mysql.CreatePost(post)
	if err != nil {
		log.Info("创建文章失败")
		return err
	}
	/*community, err := mysql.GetCommunityNameByID(strconv.FormatUint(post.CommunityID, 10))
	if err != nil {
		log.Info("获取社区id失败")
		return err
	}*/
	// redis存储帖子信息
	/*	redis.CreatePost(
		fmt.Sprint(post.PostID),
		fmt.Sprint(post.AuthorID),
		post.Title,
		)
	*/
	return nil
}
