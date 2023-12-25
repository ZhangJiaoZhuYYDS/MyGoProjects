// @Author zhangjiaozhu 2023/6/4 13:17:00
package mysql

import (
	"crypto/md5"
	"encoding/hex"
)

/*密码加密*/

func EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte("密钥"))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
