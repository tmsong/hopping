/**
剑指 Offer 12. 矩阵中的路径

给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true

示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false

思路：经典的递归回溯
*/
package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCEF"))
}

func exist(board [][]byte, word string) bool {
	chars := []byte(word)
	path := make([][]bool, 0)
	for range board {
		tmp := make([]bool, len(board[0]))
		path = append(path, tmp)
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				if find(board, path, i, j, chars) {
					return true
				}
			}
		}
	}
	return false
}

func find(board [][]byte, path [][]bool, curX, curY int, chars []byte) bool {
	if len(chars) == 0 {
		return true
	}
	if curX < 0 || curX >= len(board) || curY < 0 || curY >= len(board[0]) || path[curX][curY] {
		return false
	}
	if board[curX][curY] != chars[0] {
		return false
	}
	if len(chars) == 1 {
		return true
	}
	path[curX][curY] = true
	defer func() {
		path[curX][curY] = false
	}()
	return find(board, path, curX-1, curY, chars[1:]) ||
		find(board, path, curX+1, curY, chars[1:]) ||
		find(board, path, curX, curY-1, chars[1:]) ||
		find(board, path, curX, curY+1, chars[1:])
}
