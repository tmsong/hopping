/*
25. K 个一组翻转链表

给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

思路，不断截取链表中的部分进行切分，再进行翻转。
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 { //如果只有一节或者k<=1，直接返回
		return head
	}
	cur := head
	var next *ListNode
	for cnt := k - 1; cnt > 0; cnt-- { //先找下一段需要逆转的开头
		cur = cur.Next
		if cur == nil {
			return head //到头了，不用继续了
		}
	}
	//如果成功到了这步，那么就找到了这段需要逆转的链段头
	if cur.Next != nil {
		next = reverseKGroup(cur.Next, k) //先递归调用
	}
	newHead, newTail := reverseList(head, cur)
	newTail.Next = next
	return newHead
}

func reverseList(head, tail *ListNode) (newHead, newTail *ListNode) {
	if head == tail {
		return
	}
	newTail = head
	end := tail.Next
	var prev, next *ListNode
	for head != end {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev, newTail
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	newList := reverseKGroup(head, 3)
	fmt.Println(newList.Val)
}
