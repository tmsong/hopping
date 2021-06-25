/**
剑指 Offer 36. 二叉搜索树与双向链表

输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。

为了让您更好地理解问题，以下面的二叉搜索树为例：

我们希望将这个二叉搜索树转化为双向循环链表。链表中的每个节点都有一个前驱和后继指针。对于双向循环链表，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。

下图展示了上面的二叉搜索树转化成的链表。“head” 表示指向链表中有最小元素的节点。

特别地，我们希望可以就地完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中的第一个节点的指针。

思路：二叉搜索树的中序遍历就是排序链表，因此使用递归，则root的新left是root.left子树变换后的最右节点，root的新right是root.right子树变换后的最左节点
*/
package main

import "fmt"

func main() {
	tree := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}}},
		Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 9}},
	}
	left, right := treeToDoubleList(tree)
	fmt.Println(left.Val, right.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeToDoubleList(root *TreeNode) (left, right *TreeNode) {
	left, right = root, root
	if root == nil {
		return
	}
	if root.Left != nil {
		var r *TreeNode
		left, r = treeToDoubleList(root.Left)
		root.Left, r.Right = r, root
	}
	if root.Right != nil {
		var l *TreeNode
		l, right = treeToDoubleList(root.Right)
		root.Right, l.Left = l, root
	}
	return
}
