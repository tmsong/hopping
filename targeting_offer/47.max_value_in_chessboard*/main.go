/*
*
剑指 Offer 47. 礼物的最大价值

在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

思路：DFS不行，时间复杂度太高！得用动态规划。
设f(i,j)为从期盼左上角走至单元格(i,j)的礼物最大累计价值，则易得到递推关系：f(i,j)=max[f(i,j−1),f(i−1,j)]+grid(i,j)
*/
package main

import "fmt"

func maxValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	maxValues := make([][]int, m)
	for i := 0; i < m; i++ {
		maxValues[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				maxValues[i][j] = grid[i][j]
			} else if i == 0 {
				maxValues[i][j] = maxValues[i][j-1] + grid[i][j]
			} else if j == 0 {
				maxValues[i][j] = maxValues[i-1][j] + grid[i][j]
			} else if maxValues[i-1][j] > maxValues[i][j-1] {
				maxValues[i][j] = maxValues[i-1][j] + grid[i][j]
			} else {
				maxValues[i][j] = maxValues[i][j-1] + grid[i][j]
			}
		}
	}
	return maxValues[m-1][n-1]
}

func main() {

	grid := [][]int{
		{1, 2, 5},
		{3, 2, 1},
	}
	fmt.Println(maxValue(grid))
}
