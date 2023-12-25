// @Author zhangjiaozhu 2023/12/19 9:54:00
package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 虚拟头节点，用于结果返回
	pre := &ListNode{Val: 0}
	// 修改当前遍历位置
	cur := pre
	carry := 0 // 进位值
	// 循环判断两链表
	for l1 != nil || l2 != nil {
		// 初始化遍历当前链表对应的值
		x := 0
		y := 0

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}

		cur = cur.Next

	}
	if carry == 1 {
		cur.Next = &ListNode{Val: carry}
	}
	return pre.Next
}

// 自写

func demo(l1, l2 *ListNode) *ListNode {
	pre := &ListNode{Val: 0}
	cur := pre
	carry := 0

	for l1 != nil || l2 != nil {
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = sum / 10
		cur.Next = &ListNode{
			Val: sum % 10,
		}
		cur = cur.Next
	}
	if carry == 1 {
		cur.Next = &ListNode{Val: carry}
	}
	return pre.Next
}