/**
剑指 Offer 25. 合并两个排序的链表

输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

思路：类似于归并排序合并
*/
package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	res := mergeTwoLists(l1, l2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tmp *ListNode
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			if head == nil {
				head = l1
			}
			if tmp != nil {
				tmp.Next = l1
			}
			tmp = l1
			l1 = l1.Next
		} else {
			if head == nil {
				head = l2
			}
			if tmp != nil {
				tmp.Next = l2
			}
			tmp = l2
			l2 = l2.Next
		}
	}
	if l1 == nil {
		tmp.Next = l2
	} else if l2 == nil {
		tmp.Next = l1
	}
	return head
}
