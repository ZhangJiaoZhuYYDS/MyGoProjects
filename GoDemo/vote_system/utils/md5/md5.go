// @Author zhangjiaozhu 2023/3/16 16:40:00
package md5

import (
	"crypto/md5"
	"fmt"
)

func GetMd5(s string) string {
	return fmt.Sprintf("%x",md5.Sum([]byte(s)))
}
