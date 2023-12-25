// @Author zhangjiaozhu 2023/12/23 10:28:00
package mid

import (
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	// 去空格
	str = strings.TrimSpace(str)
	n := len(str)
	// 符号
	sign := 1
	// 结果
	res := 0
	// 字符串索引
	index := 0
	// 处理+-符号
	if str[index] == '+' || str[index] == '-' {
		if str[index] == '-' {
			sign = -1
		}
		index++
	}
	//处理数字
	for index < n && isDigit(str[index]) {
		//将字符转换成数字
		digit, _ := strconv.Atoi(string(str[index]))
		// 判断是否溢出
		if res > (1<<31-1-digit)/10 {
			if sign == 1 {
				return 1<<31 - 1
			} else {
				return -1 << 31
			}
		}

		res = res*10 + digit
		index++
	}

	return res * sign
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
