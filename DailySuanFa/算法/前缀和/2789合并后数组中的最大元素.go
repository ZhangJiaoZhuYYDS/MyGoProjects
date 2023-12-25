// @Author zhangjiaozhu 2023/12/21 14:57:00
package 前缀和

// 贪心 前缀和
func maxArrayValue(nums []int) int64 {
	count := len(nums)
	index := count - 1
	max := nums[index]
	for index > 0 {
		if max >= nums[index-1] {
			max += nums[index-1]
		} else {
			max = nums[index-1]
		}
		index--
	}
	return int64(max)
}
