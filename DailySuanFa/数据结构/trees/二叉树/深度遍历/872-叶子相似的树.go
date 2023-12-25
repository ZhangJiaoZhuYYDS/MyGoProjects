// @Author zhangjiaozhu 2023/7/28 8:57:00
package 深度遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(root *TreeNode)

	res := make([]int, 0)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 把叶子节点的值添加到切片
		if root.Left == nil && root.Right == nil {
			res = append(res, root.Val)
		}
		// 递归左子树
		dfs(root.Left)
		// 递归右子树
		dfs(root.Right)
	}
	// 遍历root1
	dfs(root1)
	// 复制res的值到res1
	res1 := append([]int{}, res...)
	// 清空res,来存储root2的叶子节点
	res = []int{}
	// 遍历root2
	dfs(root2)
	if len(res) != len(res1) {
		return false
	}
	// 比较叶子节点序列
	for i, v := range res {
		if v != res1[i] {
			return false
		}
	}
	return true
}
