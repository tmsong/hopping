/**
剑指 Offer 32 - III. 从上到下打印二叉树 III
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

例如:
给定二叉树:[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]

思路：还是递归的时候记录一下层级，如果是偶数就从右往左添加，否则从左往右添加

标准思路：双端队列，奇数层添加到尾部，偶数层添加到头部
*/
package main

import "fmt"

func main() {
	tree := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 8}, Right: &TreeNode{Val: 11}},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	res := levelOrder(tree)
	for _, s := range res {
		fmt.Println(s)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	reCur(root, 1, &res)
	return res
}

func reCur(root *TreeNode, curLevel int, res *[][]int) {
	for len(*res) < curLevel {
		*res = append(*res, make([]int, 0))
	}
	if curLevel%2 == 1 {
		(*res)[curLevel-1] = append((*res)[curLevel-1], root.Val)
	} else {
		(*res)[curLevel-1] = append([]int{root.Val}, (*res)[curLevel-1]...)
	}
	if root.Left != nil {
		reCur(root.Left, curLevel+1, res)
	}
	if root.Right != nil {
		reCur(root.Right, curLevel+1, res)
	}
}
