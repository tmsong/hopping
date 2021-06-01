/**
剑指 Offer 17. 打印从1到最大的n位数

输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:

输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]

思路：没什么可讲的，找到最大的n位数，然后循环就可以了
*/
package main

import "fmt"

func main() {
	fmt.Println(printNumbers(1))
}

func printNumbers(n int) []int {
	ret := make([]int, 0)
	for i := 1; i <= maxN(n); i++ {
		ret = append(ret, i)
	}
	return ret
}

func maxN(n int) int {
	ret := 1
	for ; n > 0; n-- {
		ret *= 10
	}
	return ret - 1
}
