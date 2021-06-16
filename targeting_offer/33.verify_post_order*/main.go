/**
剑指 Offer 33. 二叉搜索树的后序遍历序列

输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。

参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true

思路：二叉搜索树的所有左子节点，均小于根，再小于右，因此后序遍历时，可以先找到根节点（最后一个值），
然后从左向右找到第一个大于根节点的值(idx)，依据此值切开[:idx],[idx:]，左边的值都应该小于根，右边的值应该都大于根。然后递归验证
*/
package main

import "fmt"

func main() {
	fmt.Println(verifyPostorder([]int{3, 10, 6, 9, 2}))
}

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	l := len(postorder)
	//取根节点
	root := postorder[l-1]
	//去掉根节点
	postorder = postorder[:l-1]
	cut := -1 //找第一个大于root的值
	for idx, value := range postorder {
		if value > root && cut == -1 {
			cut = idx
		}
		if cut >= 0 && value < root { //说明[idx:]中有小于根的值，错误
			return false
		}
	}
	if cut == -1 { //没找到大于根的，则验证
		return verifyPostorder(postorder)
	}
	return verifyPostorder(postorder[:cut]) && verifyPostorder(postorder[cut:])
}
