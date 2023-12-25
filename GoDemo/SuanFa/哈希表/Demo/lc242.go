package main

import "fmt"

func main() {
	fmt.Println('v' - 'a')
	for i := range "sgs" {
		fmt.Println(i)
	}
	s := "anagram"
	t := "nagaram"
	println(isAnagram(s, t))
}
func isAnagram(s string, t string) bool {
	// 声明一个26大小的数组
	count := [26]int{}
	for _, v := range s {
		count[v-'a']++
	}
	for _, v := range t {
		count[v-'a']--
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	//return count == [26]int{}
	return true
}
