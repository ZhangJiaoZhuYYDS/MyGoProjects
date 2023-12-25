package main

func main() {

	nums := []int{1, 7, 4, 9, 2, 5}
	//nums := []int{1, 2, 3, 4, 5, 6}
	println(wiggleMaxLength(nums))
}

// 贪心
/*
	func wiggleMaxLength(nums []int) int {
		// 切片元素小于2 ，直接返回切片元素个数
		if len(nums) < 2 {
			return len(nums)
		}
		// 默认序列最右边有一个
		result := 1
		// 计算第一个和第二个的差值
		prediff := nums[1] - nums[0]
		// 不等于0 ，即代表开始两元素不是单调递增的
		if prediff != 0 {
			result = 2
		}
		// 遍历元素
		for i := 2; i < len(nums); i++ {
			curdiff := nums[i] - nums[i-1]
			if (curdiff > 0 && prediff <= 0) || (curdiff < 0 && prediff >= 0) {
				result++
				prediff = curdiff
			}
		}
		return result
	}
*/
// 动态规划
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up, down := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	return max(up, down)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
