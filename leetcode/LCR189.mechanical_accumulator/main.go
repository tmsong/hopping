/*
*

LCR 189. 设计机械累加器
请设计一个机械累加器，计算从 1、2... 一直累加到目标数值 target 的总和。
注意这是一个只能进行加法操作的程序，不具备乘除、if-else、switch-case、for 循环、while 循环，及条件判断语句等高级功能。

思路：不能用乘除和循环，只能用加法，可以用递归实现。用一个闭包来操作外部变量。
*/
package main

import "fmt"

func mechanicalAccumulator(target int) int {
	ans := 0
	var sum func(int) bool
	sum = func(num int) bool {
		ans += num
		return num > 0 && sum(num-1)
	}
	sum(target)
	return ans
}

func main() {
	fmt.Println(mechanicalAccumulator(10))
}
