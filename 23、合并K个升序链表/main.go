package main

import "fmt"

/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
*/

// Definition for singly-linked list.
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
					Val:  8,
					Next: nil,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  7,
					Next: nil,
				},
			},
		},
	}

	l3 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 10,
			Next: &ListNode{
				Val:  11,
				Next: nil,
			},
		},
	}
	lists := []*ListNode{l2, l1, l3}
	result := mergeKLists(lists)
	fmt.Println(result)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	return divide(lists, 0, len(lists)-1)
}

func divide(lists []*ListNode, start, end int) *ListNode {
	//先拆分
	if start == end {
		return lists[start]
	}
	mid := start + (end-start)/2
	left := divide(lists, start, mid)
	right := divide(lists, mid+1, end)
	return merge(left, right)
}

func merge(a *ListNode, b *ListNode) *ListNode {
	//递归思想  合并两个升序链表
	if a == nil {
		//a 链表为空
		return b
	}
	if b == nil {
		//b 链表为空
		return a
	}
	if a.Val < b.Val {
		a.Next = merge(a.Next, b)
		return a
	} else {
		b.Next = merge(a, b.Next)
		return b
	}
}
