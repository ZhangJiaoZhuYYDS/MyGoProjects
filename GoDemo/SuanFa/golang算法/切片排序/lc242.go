package main

import "sort"

// 242 有效的字母异位词

func main() {

}
func isAnagram(s, t string) bool {
	// 将字符串转为字节数组
	s1, s2 := []byte(s), []byte(t)
	// 对s1降序排序
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] > s1[j]
	})
	// 对s2降序排序
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] > s2[j]
	})
	// 将字节数组转为字符串
	return string(s1) == string(s2)
}
