// @Author zhangjiaozhu 2023/7/30 9:59:00
package 层序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 通过更新深度,递归传入右子树的右子树....深度从0开始，正好和数组的长度一致，然后把节点值添加到数组中去
func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0)
	var bfs func(root *TreeNode, a int)
	bfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth == len(ret) {
			ret = append(ret, root.Val)
		}
		bfs(root.Right, depth+1)
		bfs(root.Left, depth+1)
	}
	bfs(root, 0)
	return ret
}
func rightSideView2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if i == n-1 {
				res = append(res, queue[i].Val)
			}
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
