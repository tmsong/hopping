/*
*
剑指 Offer 55 - I. 二叉树的深度

输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

例如：

给定二叉树 [3,9,20,null,null,15,7]，

	  3
	 / \
	9  20
	  /  \
	 15   7

返回它的最大深度 3 。

思路：很简单，递归。
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if leftDepth, rightDepth := maxDepth(root.Left), maxDepth(root.Right); leftDepth < rightDepth {
		return rightDepth + 1
	} else {
		return leftDepth + 1
	}
}

func main() {

}
