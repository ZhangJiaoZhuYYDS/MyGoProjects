// @Author zhangjiaozhu 2023/12/19 10:28:00
package 深度遍历

// 前序遍历
func preorderTravesal(root *TreeNode) (res []int) {
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		travel(root.Left)
		travel(root.Right)
	}
	travel(root)
	return res
}

// 中序遍历
func inorderTravesal(root *TreeNode) (res []int) {
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		res = append(res, root.Val)
		travel(root.Right)
	}
	travel(root)
	return
}

// 后序遍历
func postorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}
