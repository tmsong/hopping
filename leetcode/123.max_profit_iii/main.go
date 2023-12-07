/*
*
122. 买卖股票的最佳时机 III

给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

思路：还是动态规划，
由于我们最多可以完成两笔交易，因此在任意一天结束之后，我们会处于以下五个状态中的一种：
1.未进行过任何操作；
2.只进行过一次买操作；
3.进行了一次买操作和一次卖操作，即完成了一笔交易；
4.在完成了一笔交易的前提下，进行了第二次买操作；
5.完成了全部两笔交易。
由于第一个状态的利润显然为 0，因此我们可以不用将其记录。
对于剩下的四个状态，我们分别将它们的最大利润记为buy1，sell1，buy2，sell2

则有递推关系
buy1[i]=max{buy1[i-1], −prices[i]}	//一个是不操作，一个是操作
sell1[i]=max{sell1[i-1], buy1[i-1]+prices[i]}
buy2[i]=max{buy2[i-1], -prices[i]}
sell2[i]=max{sell2[i-1], buy2[i-1]+prices[i]}

边界条件：
buy1[0] = -prices[i]
sell1[0] = 0
buy2[0] = -prices[i]
sell2[0] = 0

最终答案为sell2
*/
package main

import "fmt"

func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}
