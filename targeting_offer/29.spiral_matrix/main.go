/**
剑指 Offer 29. 顺时针打印矩阵

输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。


示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix =[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

思路：其实是考察细心的题，就像剥洋葱一样，按照层次一层一层，一圈一圈剥开。
*/
package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		[]int{2, 5, 8},
		[]int{4, 0, -1},
	}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	ret := make([]int, 0)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ret
	}
	rows := len(matrix)
	cols := len(matrix[0])
	var maxDepth int
	if rows < cols {
		maxDepth = (rows + 1) / 2
	} else {
		maxDepth = (cols + 1) / 2
	}
	for i := 1; i <= maxDepth; i++ {
		ret = append(ret, byDepth(matrix, i)...)
	}
	return ret
}

func byDepth(matrix [][]int, depth int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	ret := make([]int, 0)
	for i := depth - 1; i <= cols-depth; i++ { //横着走
		ret = append(ret, matrix[depth-1][i])
	}
	for i := depth; i <= rows-depth; i++ { //竖着走
		ret = append(ret, matrix[i][cols-depth])
	}
	if rows-depth > depth-1 {
		for i := cols - depth - 1; i >= depth-1; i-- { //返回来横着走
			ret = append(ret, matrix[rows-depth][i])
		}
	}
	if cols-depth > depth-1 {
		for i := rows - depth - 1; i >= depth; i-- { //返回来竖着走
			ret = append(ret, matrix[i][depth-1])
		}
	}
	return ret
}
