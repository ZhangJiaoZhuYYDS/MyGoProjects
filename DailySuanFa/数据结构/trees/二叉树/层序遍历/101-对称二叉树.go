// @Author zhangjiaozhu 2023/8/14 20:59:00
package 层序遍历

// 使用切片模拟队列，实现先进先出
func isSymmetric(root *TreeNode) bool {
	u, v := root, root

	// 队列
	q := []*TreeNode{}

	// 将头节点入队
	q = append(q, u)
	q = append(q, v)
	// 遍历队列
	for len(q) > 0 {

		// 每次取队列前两个元素比较
		u, v = q[0], q[1]

		// 出队
		q = q[2:]

		// 两个元素都是空，进行下一轮
		if u == nil && v == nil {
			continue
		}
		// 两个元素一个空，一个不为空，不符合
		if u == nil || v == nil {
			return false
		}
		// 两个元素的值不相等，不符合
		if u.Val != v.Val {
			return false
		}

		// 入队
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}