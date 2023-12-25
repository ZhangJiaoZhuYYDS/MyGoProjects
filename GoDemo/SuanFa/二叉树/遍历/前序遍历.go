package main

func main() {
	root1 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root2 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	root := TreeNode{
		Val:   3,
		Left:  &root1,
		Right: &root2,
	}

	preorderTraversal(&root)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}
