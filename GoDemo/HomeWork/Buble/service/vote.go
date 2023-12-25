// @Author zhangjiaozhu 2023/3/7 21:19:00
package service

import (
	"Buble/models"
	"github.com/lexkong/log"
)

func VoteForPost(useid uint64,p *models.VoteDataForm)  error{
	log.Info(string(useid),p.PostID,p.Direction)
	//return redis.VoteForPost(strconv.Itoa(int(userId)), p.PostID, float64(p.Direction))
	return nil
}
