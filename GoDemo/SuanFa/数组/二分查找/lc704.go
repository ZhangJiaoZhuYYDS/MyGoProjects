package main

/*704 二分查找*/
// 前提是排好序
func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	taget := 9
	println(search(nums, taget))
}
func search(nums []int, target int) int {
	// 获取切片长度
	numsLen := len(nums)
	// 初始化左,右指针位置
	l, r := 0, numsLen-1
	for l <= r {
		// 初始化中间位置
		m := (r-l)/2 + l
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1

}
