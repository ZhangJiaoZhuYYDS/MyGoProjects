package main

import "sort"

// 49
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	println(groupAnagrams(strs))
}

// 一 排序分类后添加到map
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	// 遍历每个字符串，将字符串转为字节切片并排序，然后把排序的字符串作为map的索引，最后把同索引的字符串添加到一块
	for _, v := range strs {
		// 将字符串转为字节切片
		s := []byte(v)
		sort.Slice(s, func(i, j int) bool {
			return s[i] > s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], v)
	}
	// 声明一个二维切片
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

// 二 哈希表计数作为map的键
func groupAnagramsHash(strs []string) [][]string {
	// 声明一个以26大小数组为键，字符串切片为值的map
	mp := map[[26]int][]string{}
	for _, str := range strs {
		count := [26]int{}
		for _, v := range str {
			count[v-'a']++
		}
		// 把相同键的值添加到同一键下的字符串切片
		mp[count] = append(mp[count], str)
	}
	//ans := [][]string{}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
