/**
剑指 Offer 06. 从尾到头打印链表

输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]

思路1：
顺便写个反转链表吧。

思路2：
递归。
*/
package main

import "fmt"

func main() {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: nil}}}
	fmt.Println(reversePrint(list))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//方法1 反转链表
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	reversed := reverseList(head)
	ret := make([]int, 0)
	for reversed != nil {
		ret = append(ret, reversed.Val)
		reversed = reversed.Next
	}
	return ret
}

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

//方法2 递归
func reversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}
	ret := reversePrint(head.Next)
	ret = append(ret, head.Val)
	return ret
}
