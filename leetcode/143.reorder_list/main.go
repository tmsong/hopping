/*
*
143. 重排链表
中等
相关标签
相关企业
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

思路：
链表先拆两半，倒序后一半，然后再拼接起来。
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow, fast := head, head
	//step 1 先找链表中点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//把slow后面的链表断开
	newHead := slow.Next
	slow.Next = nil
	newHead = reverseList(newHead)
	mergeList(head, newHead)
	return
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, next *ListNode
	for prev, next = nil, head.Next; head != nil; {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func mergeList(h1, h2 *ListNode) {
	for h2 != nil {
		h1n := h1.Next
		h1.Next = h2
		h2 = h2.Next
		h1.Next.Next = h1n
		h1 = h1n
	}
}

func main() {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	reorderList(l)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
