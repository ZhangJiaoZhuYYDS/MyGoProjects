package main

import "sort"

func main() {

}
func findContentChildren(g []int, s []int) int {
	// 排序切片
	sort.Ints(g)
	sort.Ints(s)
	// 胃口指针起始位置
	gStart := 0
	// 结果
	count := 0
	// 遍历饼干，优先小饼干满足小胃口
	for i := 0; i < len(s); i++ {
		// 饼干满足胃口，结果+1，胃口指针+1
		if s[i] >= g[gStart] {
			gStart++
			count++
		}
	}
	return count
}
