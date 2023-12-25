// @Author zhangjiaozhu 2023/9/16 20:27:00
package main

// 3 从尾到头打印链表

// ListNode /**
type ListNode struct {
	Val  int
	Next *ListNode
}

// 3.1 遍历添加到切片，然后反转切片
func reversePrint(head *ListNode) []int {
	var stack []int
	for head := head; head != nil; head = head.Next {
		stack = append(stack, head.Val)
		//head = head.Next
	}
	var ans []int
	for i := len(stack) - 1; i >= 0; i-- {
		ans = append(ans, stack[i])
	}
	return ans
}

// 使用栈结构存，然后取实现反转

// 递归
func reversePrint2(head *ListNode) []int {
	var f func(n *ListNode) []int
	f = func(n *ListNode) []int {
		if n == nil {
			return []int{}
		}
		return append(f(n.Next), n.Val)
	}
	return f(head)
}
