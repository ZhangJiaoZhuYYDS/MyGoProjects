// @Author zhangjiaozhu 2023/12/22 9:57:00
package 子序列

import "sort"

// 1 动态规划
func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	pre := getLISArray(nums)
	suf := getLISArray(reversed(nums))
	suf = reversed(suf)

	ans := 0
	for i := 0; i < n; i++ {
		if pre[i] > 1 && suf[i] > 1 {
			ans = max(ans, pre[i]+suf[i]-1)
		}
	}
	return n - ans
}

// 获取数组nums[i]的最长子序列长度
func getLISArray(nums []int) []int {
	n := len(nums)
	// dp[i]表示 在数组的前缀部分 nums[0..i]找出一个严格递增的子序列，并且以 nums[i]结束，对应着「山形状数组」的上升部分，也即最长子序列长度
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp
}

// 数组反转
func reversed(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = nums[n-1-i]
	}
	return ans
}

// 2 二分查找
func minimumMountainRemovals2(nums []int) int {
	n := len(nums)
	pre := make([]int, n)
	var g []int
	for i := 0; i < n; i++ {
		j := sort.Search(len(g), func(j int) bool {
			return g[j] >= nums[i]
		})

		if j == len(g) {
			g = append(g, nums[i])
		} else {
			g[j] = nums[i]
		}

		pre[i] = j + 1
	}

	ans := 0
	g = g[:0]
	for i := n - 1; i >= 0; i-- {
		j := sort.Search(len(g), func(j int) bool {
			return g[j] >= nums[i]
		})

		if j == len(g) {
			g = append(g, nums[i])
		} else {
			g[j] = nums[i]
		}

		suf := j + 1
		if pre[i] >= 2 && suf >= 2 {
			ans = max(ans, pre[i]+suf-1)
		}
	}

	return len(nums) - ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 3 二分查找
func minimumMountainRemovals3(nums []int) int {
	n := len(nums)
	pre := getLISArray(nums)
	suf := getLISArray(reversed(nums))
	suf = reversed(suf)

	ans := 0
	for i := 0; i < n; i++ {
		if pre[i] > 1 && suf[i] > 1 {
			ans = max(ans, pre[i]+suf[i]-1)
		}
	}
	return n - ans
}

func getLISArray2(nums []int) []int {
	n := len(nums)
	dp := make([]int, n)
	var seq []int
	for i := 0; i < n; i++ {
		// 寻找指定数字在数组中的插入位置
		it := sort.SearchInts(seq, nums[i])
		// 如果指定数字未出现过，也就是等于len(seq)就直接放入seq,并保存dp[i]
		if it == len(seq) {
			seq = append(seq, nums[i])
			dp[i] = len(seq)
		} else {
			// 指定数字出现在数组中，就修改指定位置数字，并统计dp[i]
			seq[it] = nums[i]
			dp[i] = it + 1
		}
	}
	return dp
}

func reversed2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = nums[n-1-i]
	}
	return ans
}