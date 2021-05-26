/**
剑指 Offer 13. 机器人的运动范围

地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

思路：经典的递归回溯
*/
package main

import "fmt"

func main() {
	fmt.Println(movingCount(3, 1, 0))
}

func movingCount(m int, n int, k int) int {
	path := make([][]bool, 0)
	for i := 0; i < m; i++ {
		path = append(path, make([]bool, n))
	}
	move(0, 0, m, n, path, k)
	var count int
	for i := range path {
		for j := range path[i] {
			if path[i][j] {
				count++
			}
		}
	}
	return count
}

func move(curX, curY, m, n int, path [][]bool, k int) {
	if curX < 0 || curX >= m || curY < 0 || curY >= n || path[curX][curY] {
		return
	}
	if calSum(curX, curY) <= k {
		path[curX][curY] = true
		move(curX-1, curY, m, n, path, k)
		move(curX, curY-1, m, n, path, k)
		move(curX+1, curY, m, n, path, k)
		move(curX, curY+1, m, n, path, k)
	}
}

func calSum(curX, curY int) (sum int) {
	for curX > 0 {
		sum += curX % 10
		curX /= 10
	}
	for curY > 0 {
		sum += curY % 10
		curY /= 10
	}
	return sum
}
