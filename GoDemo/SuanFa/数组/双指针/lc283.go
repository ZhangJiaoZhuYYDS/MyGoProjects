package main

import "fmt"

// 283 移动零
func main() {
	arr := []int{1}
	moveZeroes(arr)
	fmt.Print(arr)
}
func moveZeroes(nums []int) {
	// 判断数组边界
	if nums == nil || len(nums) == 0 {
		return
	}
	// 一个元素且不为0的情况直接返回
	if len(nums) == 1 {
		return
	}
	// 初始化左右指针
	l, r := 0, 0
	// 遍历数组内每个元素
	for r = 0; r < len(nums); r++ {
		// 判断当前遍历元素是否不等于0；不等于0就交换左指针和当前遍历元素,然后左指针右移一位
		if nums[r] != 0 {
			temp := nums[r]
			nums[r] = nums[l]
			nums[l] = temp
			l++
		}
	}

}
