/**
剑指 Offer 32 - I. 从上到下打印二叉树

从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。


例如:
给定二叉树:[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]

思路：二叉树的层序遍历，利用一个队列，每次从队列里取一个节点：输出此节点的值，再push此节点的左/右孩子到队列中，循环至队列为空。

*/
package main

import "fmt"

func main() {
	tree := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 8}, Right: &TreeNode{Val: 11}},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	res := levelOrder(tree)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		res = append(res, queue[0].Val)
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
	}
	return res
}
