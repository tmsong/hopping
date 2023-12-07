package main

import "fmt"

/*
将一个 二叉搜索树 就地转化为一个 已排序的双向循环链表 。

对于双向循环列表，你可以将左右孩子指针作为双向循环链表的前驱和后继指针，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。

特别地，我们希望可以 就地 完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中最小元素的指针。

思路：递归处理左子树和右子树，然后接在一起。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeToDoublyList(root *TreeNode) (left, right *TreeNode) {
	left, right = root, root
	if root.Left == nil && root.Right == nil { //叶子节点，则直接返回
		return
	}
	l, r := root.Left, root.Right
	if root.Left != nil {
		ll, lr := treeToDoublyList(l) //递归处理左子结点，然后root直接接在左链的最右。
		lr.Right, root.Left = root, lr
		left = ll //左节点返回ll
	}
	if root.Right != nil {
		rl, rr := treeToDoublyList(r)
		rl.Left, root.Right = root, rl
		right = rr
	}
	return
}
func main() {
	root := &TreeNode{Val: 6,
		Left: &TreeNode{Val: 4,
			Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 9,
			Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 10}}}
	left, right := treeToDoublyList(root)
	fmt.Println(left.Val, right.Val)
}
