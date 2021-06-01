/**
剑指 Offer 18. 删除链表的节点

给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点。

注意：此题对比原题有改动

示例 1:

输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为5的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
示例 2:

输入: head = [4,5,1,9], val = 1
输出: [4,5,9]
解释: 给定你链表中值为1的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

思路：用一个prev记录前一个，并记得把prev的Next置空，cur记录当前，如果cur.Val == val，则忽略当前节点，直接让cur = cur.Next，否则正常让prev.Next = cur

*/
package main

import "fmt"

func main() {
	list := &ListNode{Val: -3, Next: &ListNode{Val: 5, Next: &ListNode{Val: -99}}}
	newList := deleteNode(list, -99)
	for newList != nil {
		fmt.Println(newList.Val)
		newList = newList.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	var newHead, prev *ListNode
	for cur := head; cur != nil; cur = cur.Next {
		if prev != nil {
			prev.Next = nil
		}
		if cur.Val == val {
			continue
		}
		if newHead == nil {
			newHead = cur
		}
		if prev != nil {
			prev.Next = cur
		}
		prev = cur
	}
	return newHead
}

//原来的解法，更优雅。
//思路：一句prev.Next = cur.Next就解决了
func deleteNode2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}
	if head.Next == nil {
		return head
	}
	for prev, cur := head, head.Next; prev.Next != nil; prev, cur = cur, cur.Next {
		if cur.Val == val {
			prev.Next = cur.Next
			break
		}
	}
	return head
}
