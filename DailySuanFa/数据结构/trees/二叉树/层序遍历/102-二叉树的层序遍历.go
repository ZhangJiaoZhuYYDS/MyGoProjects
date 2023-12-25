// @Author zhangjiaozhu 2023/7/30 10:40:00
package 层序遍历

/*java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
*/

/*
class Solution {
public List<List<Integer>> levelOrder(TreeNode root) {
	// 返回结果
List<List<Integer>> res = new ArrayList<>();
	// 声明队列
Queue<TreeNode> queue = new ArrayDeque<>();
if (root != null) {
queue.add(root);
}
 // 以队列是否为空作为循环条件
while (!queue.isEmpty()) {
 // 获取队列元素个数
int n = queue.size();
 // 记录每层个数
List<Integer> level = new ArrayList<>();
 // 把每一层的元素添加到
for (int i = 0; i < n; i++) {
  // 队列出
TreeNode node = queue.poll();
 // 添加
level.add(node.val);
if (node.left != null) {
queue.add(node.left);
}
if (node.right != null) {
queue.add(node.right);
}
}
res.add(level);
}

return res;
}

}*/

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		level := []int{}
		for i := 0; i < n; i++ {
			// 获取元素
			node := queue[0]
			// 出队列
			queue = queue[1:]
			// 添加到每一层的数组中
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
