// @Author zhangjiaozhu 2023/9/22 10:40:00
package 排序

import "sort"

/*我们创建数组 nums\textit{nums}nums 的拷贝，记作数组 numsSorted\textit{numsSorted}numsSorted，并对该数组进行排序，然后我们从左向右找到第一个两数组不同的位置，即为 numsB\textit{nums}_Bnums
B
​
  的左边界。同理也可以找到 numsB\textit{nums}_Bnums
B
​
  右边界。最后我们输出 numsB\textit{nums}_Bnums
B
​
  的长度即可。
*/

func findUnsortedSubarray(nums []int) int {
	// 判断是否已经排好序了
	if sort.IntsAreSorted(nums) {
		return 0
	}
	sortArr := append([]int{}, nums...)
	sort.Sort(sort.IntSlice(sortArr))
	// sort.Ints(sortArr)
	l := 0
	r := len(nums) - 1
	// 从左边开始比较，找到最开始不一样的位置
	for nums[l] == sortArr[l] {
		l++
	}
	// 从右边开始比较，找到最后不一样的位置
	for nums[r] == sortArr[r] {
		r--
	}
	return r - l + 1
}
