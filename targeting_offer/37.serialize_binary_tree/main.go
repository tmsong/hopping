/**
剑指 Offer 37. 序列化二叉树

请实现两个函数，分别用来序列化和反序列化二叉树。

你需要设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，
你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

思路：只需要想好怎么表达节点间的关系即可，即用数组存储节点，用数组下标来表示节点的left和right
*/
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println(Serialize((*TreeNode)(nil)))
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}
	serialized := Serialize(root)
	deserialized := Deserialize(serialized)
	fmt.Println(deserialized.Val)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val   int `json:"val"`
	Left  int `json:"left"`
	Right int `json:"right"`
}

func Serialize(root *TreeNode) string {
	nodes := make([]*Node, 0)
	if root != nil {
		serialize(root, &nodes)
	}
	res, _ := json.Marshal(nodes)
	return string(res)
}

func serialize(root *TreeNode, nodes *[]*Node) (idx int) {
	idx = len(*nodes)
	node := &Node{Val: root.Val}
	*nodes = append(*nodes, node)
	if root.Left != nil {
		node.Left = serialize(root.Left, nodes)
	}
	if root.Right != nil {
		node.Right = serialize(root.Right, nodes)
	}
	return
}

func Deserialize(data string) (root *TreeNode) {
	nodes := make([]*Node, 0)
	_ = json.Unmarshal([]byte(data), &nodes)
	if len(nodes) == 0 {
		return nil
	}
	treeNodes := make([]*TreeNode, len(nodes))
	for idx, node := range nodes {
		treeNodes[idx] = &TreeNode{Val: node.Val}
	}
	for idx, node := range nodes {
		if node.Left > 0 {
			treeNodes[idx].Left = treeNodes[node.Left]
		}
		if node.Right > 0 {
			treeNodes[idx].Right = treeNodes[node.Right]
		}
	}
	return treeNodes[0]
}
