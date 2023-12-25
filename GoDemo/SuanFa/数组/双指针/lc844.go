package main

import "fmt"

// 844. 比较含退格的字符串
func main() {
	s, _ := "ab##c", "ad#c"
	println(test(s))
}

//func backspaceCompare(s string, t string) bool {
//
//}
func test(s *string) string {
	newS := ""
	// 遍历字符串
	for i, v := range *s {
		// 判断是否为'#'
		if v == 35 {
			newS = *s[0 : i-1]
			s = newS
		} else {
			newS += fmt.Sprintf("%c", v)
			s = newS
		}
	}
	return newS
}
