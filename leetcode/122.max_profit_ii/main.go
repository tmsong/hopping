/*
*
122. 买卖股票的最佳时机 II

给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。

思路：动态规划，用dp[i][0]表示当天不持有股票时的最大利润，用dp[i][1]表示当天持有股票时的最大利润。

则有递推关系：
dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price[i])
dp[i][1] = max(dp[i-1][0]-price[i], dp[i-1][1])

最大利润：最后一天的dp[i][0]
*/
package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	notHold, hold := make([]int, len(prices)), make([]int, len(prices))
	//初始化第一个点，然后开始递推
	notHold[0], hold[0] = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		notHold[i] = max(notHold[i-1], hold[i-1]+prices[i])
		hold[i] = max(notHold[i-1]-prices[i], hold[i-1])
	}
	return notHold[len(notHold)-1] //返回最后一个点
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
