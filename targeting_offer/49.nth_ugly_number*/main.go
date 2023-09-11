/*
*
我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:
1 是丑数。
n 不超过1690。

思路：动态规划
定义数组 dp[]，其中 dp[i]第 i个丑数.
定义三个指针 p2,p3,p5，表示下一个丑数是当前指针指向的丑数乘以对应的质因数。
(也就是说，需要一直保证dp[px] * x 要大于当前最大的丑数。这样下一个丑数一定从这三个dp[px] * x中产生）
初始时，三个指针的值都是 1。
当 2≤i≤n时，令 dp[i]=min(dp[p2]×2,dp[p3]×3,dp[p5]×5)，
然后分别比较 dp[i] 和 dp[p2]×2,dp[p3]×3,dp[p5]×5是否相等，如果相等则将对应的指针加1。
*/
package main

import "fmt"

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	dp2, dp3, dp5 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = dp[dp2] * 2
		if dp[i] > dp[dp3]*3 {
			dp[i] = dp[dp3] * 3
		}
		if dp[i] > dp[dp5]*5 {
			dp[i] = dp[dp5] * 5
		}
		if dp[i] == dp[dp2]*2 {
			dp2++
		}
		if dp[i] == dp[dp3]*3 {
			dp3++
		}
		if dp[i] == dp[dp5]*5 {
			dp5++
		}
	}
	return dp[n-1]
}
func main() {
	fmt.Println(nthUglyNumber(10))
}
