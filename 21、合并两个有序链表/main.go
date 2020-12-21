package main

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
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
	println(mergeTwoLists(l1, l2))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		head, tail, l1 = l1, l1, l1.Next
	} else {
		head, tail, l2 = l2, l2, l2.Next
	}

	//循环,直到某一个链表已遍历完
	for l1 != nil && l2 != nil {
		//找到下一个节点,添加到新链表的尾
		if l1.Val < l2.Val {
			tail.Next, l1 = l1, l1.Next
		} else {
			tail.Next, l2 = l2, l2.Next
		}
		//更新tail
		tail = tail.Next
	}

	//剩下的节点字节拼接到新链表尾部
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}

	return head
}
