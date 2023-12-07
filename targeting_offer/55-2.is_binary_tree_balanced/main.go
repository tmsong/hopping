/*
*
剑指 Offer 55 - II. 平衡二叉树

输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

	  3
	 / \
	9  20
	  /  \
	 15   7

返回 true 。

思路：很简单，还是递归
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	b, _ := isBalancedAndDepth(root)
	return b
}

func isBalancedAndDepth(n *TreeNode) (isBalanced bool, depth int) {
	if n == nil {
		return true, 0
	}
	if leftBalanced, leftDepth := isBalancedAndDepth(n.Left); !leftBalanced { //分别统计两边是否平衡&两边的深度，然后对比
		return false, 0
	} else if rightBalanced, rightDepth := isBalancedAndDepth(n.Right); !rightBalanced {
		return false, 0
	} else if leftDepth >= rightDepth && leftDepth-rightDepth <= 1 {
		return true, leftDepth + 1
	} else if rightDepth > leftDepth && rightDepth-leftDepth == 1 {
		return true, rightDepth + 1
	}
	return false, 0
}

func main() {

}
