// @Author zhangjiaozhu 2023/3/14 15:08:00
package dao

import "TouPiaoSystem/model"

// 判断用户是否投过票 通过id获取投票状态
func FindUserStatus( id uint64 , tp string) int {
	var opt model.Votesopt
	model.DB.Self.Debug().Model(&model.Votesopt{}).Where("user_id = ? AND type = ?" , id,tp).First(&opt)
	return opt.Status
}
