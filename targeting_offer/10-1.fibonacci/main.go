/**
剑指 Offer 10- I 斐波那契数列

写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：

F(0) = 0, F(1)= 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

思路：就是个循环
*/
package main

import "fmt"

func main() {
	fmt.Println(fib(5))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	n_2, n_1 := 0, 1
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
