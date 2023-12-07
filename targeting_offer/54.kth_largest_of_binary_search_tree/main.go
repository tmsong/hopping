/*
*
剑指 Offer 54. 二叉搜索树的第k大节点

给定一棵二叉搜索树，请找出其中第 k 大的节点的值。

示例 1:

输入: root = [3,1,4,null,2], k = 1

	  3
	 / \
	1   4
	 \
	  2

输出: 4

思路：二叉搜索树的中序遍历是排序数组，所以利用反的中序遍历（右->中->左），边遍历边统计个数得到第k个。
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	return *midOrderByK(root, &k)
}

func midOrderByK(n *TreeNode, k *int) *int {
	if n == nil {
		return nil
	}
	if ret := midOrderByK(n.Right, k); ret != nil {
		return ret
	}
	*k = *k - 1
	if *k == 0 { //说明这就是中序遍历的第k个节点。
		return &n.Val
	}
	return midOrderByK(n.Left, k)
}
func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 1}},
			Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6}}
	fmt.Println(kthLargest(root, 3))
}
