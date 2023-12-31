// @Author zhangjiaozhu 2023/3/7 17:17:00
package models

import "time"

type Community struct {
	CommunityID uint64 `json:"community_id" db:"community_id"`
	CommunityName string `json:"community_name" db:"community_name"`
}

type CommunityDetail struct {
	CommunityID uint64 `json:"community_id" db:"community_id"`
	CommunityName string `json:"community_name" db:"community_name"`
	Introduction string `json:"introduction,omitempty" db:"introduction"`  // omitempty 当introduction为空时，不展示
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
