// @Author zhangjiaozhu 2023/3/31 17:36:00
package main

type ListNode struct {
	Val int
	Next *ListNode
}
// 1 递归 返回尾部节点
func reverseList(head *ListNode) *ListNode {
	// 递归判断结束条件，返回尾部节点
	if head == nil || head.Next == nil {
		return head
	}
	// 回溯获取尾部节点
	newHead := reverseList(head.Next)
	// 回溯交换元素位置
	head.Next.Next = head
	head.Next = nil
	// 回溯返回尾部节点
	return newHead
}
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 获取尾部节点
	newHead := reverseList2(head.Next)
	newHead.Next = head
	head.Next = nil
	return head
}
// 2 迭代
func reverseList3(head *ListNode) *ListNode{
	var preNode *ListNode
	curNode := head
	for curNode != nil {
		tmp := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = tmp
	}
	return preNode
}

func main()  {
	head2 := ListNode{
		Val:  2,
		Next: nil,
	}
	head1 := ListNode{
		Val:  1,
		Next: &head2,
	}
	head := ListNode{
		Val:  0,
		Next: &head1,
	}
	node := reverseList2(&head)
	print(node.Val)
}