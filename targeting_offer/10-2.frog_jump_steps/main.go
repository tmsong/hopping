/**
剑指 Offer 10- II 青蛙跳台阶问题

一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

思路：与斐波那契数列数列一样，只不过a0=1,a1=1
*/
package main

import "fmt"

func main() {
	fmt.Println(numWays(7))
}

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	n_2, n_1 := 1, 1
	var ret int
	fibonacci := make([]int, n+1)
	fibonacci[0], fibonacci[1] = 0, 1
	for i := 2; i <= n; i++ {
		ret = (n_2 + n_1) % (1e9 + 7)
		n_2 = n_1
		n_1 = ret
	}
	return ret
}
