/**
剑指 Offer 26. 树的子结构
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

  3
  / \
 4  5
 / \
1  2
给定的树 B：

 4
 /
1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：

输入：A = [1,2,3], B = [3,1]
输出：false

示例 2：

输入：A = [3,4,5,1,2], B = [4,1]
输出：true

思路：递归处理，遍历A，在每个A的节点递归调用对比B值的方法

*/
package main

import "fmt"

func main() {
	A := &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 5}}
	B := &TreeNode{Val: 4, Left: &TreeNode{Val: 1}}
	fmt.Println(isSubStructure(A, B))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (reCur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func reCur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	} else if A == nil {
		return false
	} else if A.Val != B.Val {
		return false
	}
	return reCur(A.Left, B.Left) && reCur(A.Right, B.Right)
}
