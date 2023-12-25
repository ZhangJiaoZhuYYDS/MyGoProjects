package main

import "sort"

func intersect1(nums1 []int, nums2 []int) []int {
	// 排序+双指针
	// 给连个数组先排序
	sort.Ints(nums1)
	sort.Ints(nums2)
	// 初始化双指针位置
	l := 0
	r := 0
	// 获取两个数组的长度
	len1 := len(nums1)
	len2 := len(nums2)

	// 初始化返回数组
	ans := []int{}
	// 添加双指针边界判断
	for l < len1 && r < len2 {
		if nums1[l] < nums2[r] {
			l++
		} else if nums1[l] > nums2[r] {
			r++
			//当双指针指向位置的元素都相同，就把该元素添加到数组里面，然后同时移动双指针
		} else {
			ans = append(ans, nums1[l])
			l++
			r++
		}
	}
	return ans
}

// 哈希表方法
func intersect(nums1 []int, nums2 []int) []int {
	// 如果num1的长度大于nums2的长度，就递归方法来调换nums1和nums2的长度
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	// 声明一个map
	m := map[int]int{}
	// 遍历nums1,map记录每个元素的次数
	for _, num := range nums1 {
		m[num]++
	}
	// 声明一个记录结果的数组
	intersection := []int{}
	// 遍历nums2 ,判断map是否含有遍历元素的键，若有就添加到数组里面，然后map元素就减一
	for _, num := range nums2 {
		if m[num] > 0 {
			intersection = append(intersection, num)
			m[num]--
		}
	}
	return intersection
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	println(intersect1(nums1, nums2)[0], intersect(nums1, nums2)[1])
}
