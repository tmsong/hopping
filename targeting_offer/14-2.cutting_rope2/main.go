/**
剑指 Offer 14- II. 剪绳子 II

给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m段（m、n都是整数，n>1并且m>1），
每段绳子的长度记为 k[0],k[1]...k[m - 1] 。请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

思路：要取模，那就只能用14-1的第二种思路了，直接按3来切
*/
package main

import "fmt"

func main() {
	fmt.Println(cuttingRope(5))
}

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	} else {
		return cut(n)
	}
}

func cut(n int) int {
	if n <= 4 {
		return n
	} else {
		return (3 * cut(n-3)) % (1e9 + 7)
	}
}
