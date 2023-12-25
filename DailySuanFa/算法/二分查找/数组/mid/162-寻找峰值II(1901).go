// @Author zhangjiaozhu 2023/12/19 14:52:00
package mid

func findPeakGrid(mat [][]int) []int {
	// 声明一个函数，用于寻找一个数组中最大值的索引
	var indexOfMax func([]int) int
	indexOfMax = func(a []int) (idx int) {
		for i, x := range a {
			if x > a[idx] {
				idx = i
			}
		}
		return
	}
	left, right := 0, len(mat)-2
	for left <= right {
		// 行
		i := left + (right-left)/2
		// 列
		j := indexOfMax(mat[i])

		if mat[i][j] > mat[i+1][j] {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return []int{left, indexOfMax(mat[left])}
}

func findPeakGrid2(mat [][]int) []int {
	n := len(mat)
	l, r, res := 0, n-1, n-1
	for l < r {
		mid := l + (r-l)/2
		j := indexOfMax(mat[mid])
		if mat[mid][j] > mat[mid+1][j] {
			r = mid
			res = r
		} else {
			l = mid + 1
		}
	}
	return []int{res, indexOfMax(mat[res])}
}
func indexOfMax(a []int) (idx int) {
	for i, x := range a {
		if x > a[idx] {
			idx = i
		}
	}
	return
}
