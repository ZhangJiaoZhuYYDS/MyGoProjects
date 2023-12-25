package main

// 383 赎金信
func main() {
	ransomNote := "aa"
	magazine := "aab"
	canConstruct(ransomNote, magazine)
}
func canConstruct(ransomNote string, magazine string) bool {
	// 26大小的数组
	count := [26]int{}
	// 遍历magazine,让数组记录每个元素出现次数
	for _, v := range magazine {
		count[v-'a']++
	}
	// 遍历ransomNote , 让数组记录的元素次数--
	for _, v := range ransomNote {
		count[v-'a']--
	}
	// 遍历数组，看是否存在负数
	for i := range count {
		if count[i] < 0 {
			return false
		}
	}
	return true
}
