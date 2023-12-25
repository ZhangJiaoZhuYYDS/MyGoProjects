package main

//209. 长度最小的子数组
func main() {
	var target int = 7
	nums := []int{2, 3, 1, 2, 4, 3}
	println(minSubArrayLen(target, nums))
}
func minSubArrayLen(target int, nums []int) int {
	//窗口左端点位置
	l := 0
	// 切片长度
	length := len(nums)
	// 子数组的和
	sum := 0
	// 结果
	result := length
	// 窗口右端点遍历窗口元素
	for r := 0; r < length; r++ {
		// 计算子数组的和
		sum += nums[r]
		// 循环判断sun的和是否
		for sum >= target {
			// 计算子数组的长度
			subLength := r - l + 1
			// 判断子数组长度是否小于结果
			if subLength < result {
				// 小于的话，就更新结果为子数组的长度
				result = subLength
			}
			// 滑动窗口的关键之处，减去窗口左端点的元素（就是下一个窗口元素的和），然后l++(就是窗口左端点左移一位，就是改变窗口的位置)
			sum -= nums[l]
			l++
		}

	}
	// 判断是否存在“不存在符合条件的子数组，返回0
	if result == length+1 {
		return 0
	} else {
		return result
	}
}
func minSubArrayLen2(target int, nums []int) int {
	// 记录最小子数组的长度
	//result := 0
	// 切片长度
	numsLen := len(nums)
	// 记录子数组的长度
	numLen := 1000
	// 记录子数组的和
	sum := 0
	for l := 0; l < numsLen; l++ {
		for r := l; r < numsLen; r++ {
			// 计算子数组的和
			sum += nums[r]
			// 判断子数组的和是否大于等于target
			if sum >= target {
				// 子数组的和大于等于target时，子数组的长度
				length := r - l + 1
				// 滑动窗口的关键步骤：减去num[l]（窗口的左端点位置），sum就是下一个窗口的起始和，然后l++(更新窗口的左端点位置)
				sum = sum - nums[l]
				l++
				if length <= numLen {
					numLen = length
				}
			}
		}
	}
	return numLen

}
