/**
剑指 Offer 35. 复杂链表的复制
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

思路1：构建一个map存node，但是需要额外空间
思路2：拼接+拆分，经典做法，原节点1->新节点1->原节点2->新节点2，然后copy random，然后再将链表拆开。
*/
package main

import "fmt"

func main() {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n1.Next, n2.Next, n3.Next = n2, n3, n4
	n1.Random, n2.Random, n3.Random, n4.Random = n2, n3, n1, n4
	newHead := copyRandomList(n1)
	for tmp := newHead; tmp != nil; tmp = tmp.Next {
		fmt.Println(tmp.Val)
	}
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	//第一步，变成 原节点1->新节点1->原节点2->新节点2 的形式
	for cur := head; cur != nil; {
		newNode := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = newNode
		cur = newNode.Next
	}
	newHead := head.Next
	//第二步，复制random指针
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	//第三步，拆分
	for cur, next := head, head.Next; next != nil; cur, next = next, next.Next {
		cur.Next = next.Next
	}
	return newHead
}
