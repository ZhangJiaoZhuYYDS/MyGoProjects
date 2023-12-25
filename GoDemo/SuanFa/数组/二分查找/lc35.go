package main

/* 35 搜索插入位置*/
func main() {
	nums := []int{1, 3, 5, 6}
	target := 0
	println(searchInsert2(nums, target))
}

// 遍历判断
func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		// 因为是有序数组，一旦发现大于或者等于目标值的元素，就返回i(要插入的位置)
		// 目标值在数组所有元素之前
		// 目标值等于数组中某一个元素
		// 目标值插入数组中的位置
		if num >= target {
			return i
		}
	}
	// 目标值在所有元素值后
	return len(nums)
}

// 二分查找方法
func searchInsert2(nums []int, target int) int {
	//初始化左右指针以及中间位置
	l, m, r := 0, 0, len(nums)-1
	for l <= r {
		m = (r-l)/2 + l
		// 1 目标值等于元素
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	// 包含2.目标值在数组所有元素之前 3.目标值插入数组中 4.目标值在数组所有元素之后 return right + 1;
	// 2.目标值在数组所有元素之前  此时for循环最后一次循环开始时，l ,r ，m都是 0，最后目标值在所有元素之前，所以 r = m-1 = -1 ,正好插入位置是 r+1 = -1 + 1 = 0
	return r + 1
}
