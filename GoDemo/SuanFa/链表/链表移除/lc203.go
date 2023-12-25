package main

// 203 移除链表元素

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// 1 模拟头节点
func removeElements(head *ListNode, val int) *ListNode {
	// 定义虚拟头节点
	dumpHead := &ListNode{}
	// 保存head
	dumpHead.Next = head
	// 定义当前节点为虚拟头节点
	cur := dumpHead
	for cur != nil && cur.Next != nil {
		// 判断当前元素是否为target
		if cur.Val == val {
			// 断开之间的节点，建立新节点
			cur.Next = cur.Next.Next
		} else {
			// 更新当前节点
			cur = cur.Next
		}
	}
	// 返回头节点
	return dumpHead.Next
}
