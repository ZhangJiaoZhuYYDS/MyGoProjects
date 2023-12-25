// @Author zhangjiaozhu 2023/12/20 17:45:00
package 二维数组

func areSimilar(mat [][]int, k int) bool {
	row, colum := len(mat), len((mat)[0]) // 行 列
	k %= colum                            // 求余
	for _, r := range mat {
		// go 1.21新用法，用来比较切片是否相等
		// 把原切片和移动后的切片进行比较
		if !slices.Equal(r, append(r[:k], r[k:]...)) {
			return false
		}
	}
	return true
}
