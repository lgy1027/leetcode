package main

import "fmt"

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。



示例：

给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
	l1 = reverseKGroup(l1, 3)
	fmt.Println(l1)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var pre, next *ListNode
	var cur, p = head, head
	groupSize, index := 0, 0
	// 将链表的前k个结点组成一组
	for p != nil && groupSize < k {
		p = p.Next
		groupSize++
	}
	if groupSize == k {
		// 翻转该k个结点组成的子链表
		for cur != nil && index < k {
			next = cur.Next
			cur.Next = pre
			pre = cur
			cur = next
			index++
		}
		// 继续递归地执行上述过程
		if next != nil {
			// head指向子链表翻转后的尾节点
			head.Next = reverseKGroup(next, k)
		}
		// pre记录了为子链表翻转后的头结点
		return pre
	} else {
		// 子链表长度不足k，不翻转
		return head
	}
}
