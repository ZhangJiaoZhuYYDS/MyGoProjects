// @Author zhangjiaozhu 2023/3/18 17:41:00
package uuid

import uuid "github.com/satori/go.uuid"

// 设置uuid
func GetUUID() string {
	return uuid.NewV4().String()
}
