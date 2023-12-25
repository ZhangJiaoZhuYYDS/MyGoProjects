// @Author zhangjiaozhu 2023/12/20 16:37:00
package mid

// 贪心
func maxArrayValue(nums []int) int64 {
	count := len(nums)
	index := count - 1
	max := nums[index]
	for index > 0 {
		if max > nums[index-1] {
			max += nums[index-1]
		} else {
			max = nums[index-1]
		}
		index--
	}
	return int64(max)
}
