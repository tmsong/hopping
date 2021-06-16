/**
剑指 Offer 32 - II. 从上到下打印二叉树 II

从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

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
  [9,20],
  [15,7]
]

思路：层序遍历的变种，可以用一个辅助队列记录当前入队的元素的层级。
当然，还有更简单的思路，可以递归遍历树时带一个层级进去
*/
package main

import "fmt"

func main() {
	tree := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 8}, Right: &TreeNode{Val: 11}},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	res := levelOrder2(tree)
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
	queue := make([]*TreeNode, 0)
	levelQueue := make([]int, 0)
	level := 1
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue = append(queue, root)
	levelQueue = append(levelQueue, level)
	for len(queue) > 0 {
		level := levelQueue[0]
		for len(res) < level { //补齐输出日志
			res = append(res, make([]int, 0))
		}
		res[level-1] = append(res[level-1], queue[0].Val)
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
			levelQueue = append(levelQueue, level+1)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
			levelQueue = append(levelQueue, level+1)
		}
		queue = queue[1:]
		levelQueue = levelQueue[1:]
	}
	return res
}

func levelOrder2(root *TreeNode) [][]int {
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
	(*res)[curLevel-1] = append((*res)[curLevel-1], root.Val)
	if root.Left != nil {
		reCur(root.Left, curLevel+1, res)
	}
	if root.Right != nil {
		reCur(root.Right, curLevel+1, res)
	}
}
