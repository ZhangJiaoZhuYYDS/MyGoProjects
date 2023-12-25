package main

// 27 移除元素
/*示例 2: 给定 nums = [0,1,2,2,3,0,4,2], val = 2, 函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。*/
func main() {
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	nums2 := []int{3, 2, 2, 3}
	val := 2 // 函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
	println(removeElement(nums2, val))
}
func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	// 初始化左右指针位置
	l, r := 0, len(nums)
	// 遍历元素是否等于目标值，并确定边界
	for l < r {
		// 如果当前遍历元素等于目标值，就交换当前元素和右指针的位置，即把目标值全部放在数组末尾
		if nums[l] == val {
			nums[l] = nums[r-1]
			// 每次右指针左移一位-1
			r--
		} else {
			// 每次左指针右移一位+1
			l++
		}
	}
	return l
}
