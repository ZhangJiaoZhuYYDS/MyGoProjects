// @Author zhangjiaozhu 2023/9/16 20:16:00
package main

import "sort"

/*二维数组中的查找    1 找规律  2 二分查找 */

// 1 二分查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, row := range matrix {
		i := sort.SearchInts(row, target)
		if i < len(row) && row[i] == target {
			return true
		}
	}
	return false
}

// 2 从左上到右下      遍历是大的  就往左找      遍历是小的 就往下找
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
