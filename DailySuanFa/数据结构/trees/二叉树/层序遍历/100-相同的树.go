// @Author zhangjiaozhu 2023/8/14 21:16:00
package 层序遍历

// 同101 使用队列实现

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	query := []*TreeNode{}

	query = append(query, p)
	query = append(query, q)

	for len(query) > 0 {
		u, v := query[0], query[1]
		query = query[2:]

		if u == nil && v == nil {
			continue
		}

		if u == nil || v == nil {
			return false
		}

		if u.Val != v.Val {
			return false
		}

		query = append(query, u.Left)
		query = append(query, v.Left)

		query = append(query, u.Right)
		query = append(query, v.Right)
	}
	return true
}
