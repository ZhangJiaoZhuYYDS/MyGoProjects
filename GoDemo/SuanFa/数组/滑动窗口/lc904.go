package main

// 904 水果成篮
func main() {
	fruits := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	println(totalFruit(fruits))
}
func totalFruit(fruits []int) int {
	// 声明map用于存储每个类型的数量
	cnt := map[int]int{}
	// 窗口左端点
	left := 0
	// 最长数量
	ans := 0
	// 遍历水果类型
	for right, v := range fruits {
		// 类型水果数量+1
		cnt[v]++
		// 当水果类型大于2时，不断左移动窗口，循环判断
		for len(cnt) > 2 {
			// 获取左端点位置元素（何种类型的水果）
			l := fruits[left]
			// 左端点类型水果数量减一
			cnt[l]--
			// 如果左端点这种类型的水果数量为0，就把这种类型的水果移除
			if cnt[l] == 0 {
				delete(cnt, l)
			}
			// 移动左端点位置，左移动窗口位置
			left++

		}
		ans = max(ans, right-left+1)
	}
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
