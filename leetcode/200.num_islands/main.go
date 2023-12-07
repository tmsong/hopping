/*
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：

输入：grid = [

	["1","1","1","1","0"],
	["1","1","0","1","0"],
	["1","1","0","0","0"],
	["0","0","0","0","0"]

]
输出：1

思路：
并查集，遍历一遍，把1相连。
*/
package main

import "fmt"

func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}
	parent := make([]int, rows*cols) //用parent数组来记录相连关系
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			pos := i*cols + j //求出在parent数组中的位置
			if grid[i][j] == '0' {
				parent[pos] = -1
			} else {
				parent[pos] = pos
			}
		}
	}
	var find func(int) int
	find = func(x int) int { //用于找到并查集的根节点
		if parent[x] != x && parent[x] != -1 {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) { //用于连接两个点
		parent[find(from)] = find(to)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			pos := i*cols + j //求出在parent数组中的位置
			if grid[i][j] == '0' {
				continue
			}
			if i+1 < rows && grid[i+1][j] == '1' { //将下边节点连接到此节点
				union(pos+cols, pos)
			}
			if j+1 < cols && grid[i][j+1] == '1' { //将右边节点连接到此节点
				union(pos+1, pos)
			}
		}
	}
	m := make(map[int]bool)
	for i, p := range parent {
		if p != -1 {
			m[find(i)] = true
		}
	}
	return len(m)
}

func main() {
	grid := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}
