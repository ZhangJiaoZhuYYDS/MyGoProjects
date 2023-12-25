// @Author zhangjiaozhu 2023/7/28 10:26:00
package mid

import "math"

// 二分查找
func findPeakElement(nums []int) int {
	n := len(nums)

	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}
	left, right := 0, n-1
	for {
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

func FindPeakElement(nums []int) int {
	maxIndex := 0
	for i, num := range nums {
		if num > nums[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}
