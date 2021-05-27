/**
剑指 Offer 14- I. 剪绳子

给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），
每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

示例 1：

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1
示例2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36

思路1：长度为 n 的绳子所能达到的最大乘积p[n] = max(k[x] * k[n-x])

思路2：数学问题，其实是对x ^ (n/x)（x为自变量）求极值，最后结论应该是按3最大。

*/
package main

import "fmt"

func main() {
	fmt.Println(cuttingRope(5))
}

//思路1：动态规划
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], dp[j]*dp[i-j])
		}
	}
	return dp[n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//思路2：直接切
func cuttingRope2(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	} else {
		return cut(n)
	}
}

func cut(n int) int {
	if n <= 4 {
		return n
	} else {
		return 3 * cut(n-3)
	}
}
