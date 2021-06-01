/**
剑指 Offer 16. 数值的整数次方

实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。

示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25

思路：分正负情况讨论，
问题：一个一个乘，性能太差，所以不过。这里可以把n转化成二进制表示，因此第i位就代表x^(2^i)
*/
package main

import "fmt"

func main() {
	fmt.Println(myPow(2.0, -2))
}

func myPow(x float64, n int) float64 {
	var ret float64 = 1
	var negative bool
	if n < 0 {
		negative = true
		n = -n
	}
	for tmp := x; n > 0; {
		if n&1 == 1 {
			ret *= tmp
		}
		tmp *= tmp
		n = n >> 1
	}
	if negative {
		ret = 1 / ret
	}
	return ret
}
