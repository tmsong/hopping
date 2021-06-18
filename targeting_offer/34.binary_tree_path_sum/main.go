/**
剑指 Offer 34. 二叉树中和为某一值的路径

输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

示例:
给定如下二叉树，以及目标和target = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]

思路：当然是DFS啊。

*/
package main

import "fmt"

func main() {
	tree := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}}
	res := pathSum(tree, 22)
	for _, path := range res {
		fmt.Println(path)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	path := make([]int, 0)
	find(root, target, path, &res)
	return res
}

func find(root *TreeNode, target int, path []int, res *[][]int) {
	if root.Left == nil && root.Right == nil {
		if target == root.Val {
			p := make([]int, len(path))
			copy(p, path)
			p = append(p, root.Val)
			*res = append(*res, p)
		}
		return
	}
	target -= root.Val
	path = append(path, root.Val)
	if root.Left != nil {
		find(root.Left, target, path, res)
	}
	if root.Right != nil {
		find(root.Right, target, path, res)
	}
}
