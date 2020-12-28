package main

import "fmt"

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

例：
输入：head = [1,2,3,4]
输出：[2,1,4,3]

示例 2：

输入：head = []
输出：[]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val:  10,
						Next: nil,
					},
				},
			},
		},
	}
	l1 = swapPairs(l1)
	fmt.Println(l1)
}

func swapPairs(head *ListNode) *ListNode {
	h := &ListNode{Next: head}
	pre, p := h, head
	for p != nil && p.Next != nil {
		pre.Next, p.Next = p.Next, p.Next.Next
		pre.Next.Next, p = p, p.Next
		pre = pre.Next.Next
	}
	return h.Next
}
