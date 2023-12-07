/*
*
剑指 Offer 52. 两个链表的第一个公共节点

输入两个链表，找出它们的第一个公共节点。

思路：先统计两个链表的长度，计算差值，然后让长的那个链表先走出这个差值。
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := lenList(headA), lenList(headB)
	if lenA == 0 || lenB == 0 { //边界判断
		return nil
	}
	var lenDiff int
	var longList, shortList *ListNode
	if lenA > lenB {
		lenDiff = lenA - lenB
		longList, shortList = headA, headB
	} else {
		lenDiff = lenB - lenA
		longList, shortList = headB, headA
	}
	for ; lenDiff > 0; lenDiff-- {
		longList = longList.Next
	}
	for longList != nil {
		if longList == shortList {
			break
		}
		longList = longList.Next
		shortList = shortList.Next
	}
	return longList
}

func lenList(head *ListNode) (len int) {
	for head != nil {
		len++
		head = head.Next
	}
	return
}

func main() {

}
