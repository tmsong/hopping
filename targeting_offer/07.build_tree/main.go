/**
剑指 Offer 07. 重建二叉树

输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

例如，给出

前序遍历 preorder =[3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

思路：
前序遍历中，树的每个节点一定出现在它的两个叶子节点的前面，利用这种思路，重建递归。

*/
package main

import "fmt"

func main() {
	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(tree.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	for preIdx := range preorder {
		for inIdx := range inorder {
			if preorder[preIdx] == inorder[inIdx] {
				node := &TreeNode{Val: preorder[preIdx]}
				node.Left = buildTree(preorder[preIdx+1:], inorder[:inIdx])
				node.Right = buildTree(preorder[preIdx+1:], inorder[inIdx+1:])
				return node
			}
		}
	}
	return nil
}
