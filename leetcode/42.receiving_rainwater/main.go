/*
*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

思路：动态规划
对于下标 i，下雨后水能到达的最大高度等于下标 i 两边的最大高度的最小值，
下标 i 处能接的雨水量等于下标 i 处的水能到达的最大高度减去 height[i]。

所以，先遍历一遍得到每个点i处左边的最大高度maxLeft[i]，再遍历一遍得到每个点i处右边的最大高度maxRight[i]。
然后再遍历一遍计算所有的 min(maxLeft[i],maxRight[i]) - height[i]
*/
package main

import "fmt"

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	maxLeft, maxRight := make([]int, len(height)), make([]int, len(height))
	//遍历一遍得到maxLeft
	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}
	//遍历一遍得到maxRight
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}
	var sum int
	for i := 0; i < len(height); i++ {
		sum += max(min(maxLeft[i], maxRight[i])-height[i], 0)
	}
	return sum
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

func main() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
