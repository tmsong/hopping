/**
剑指 Offer 15 二进制中1的个数

请实现一个函数，输入一个整数（以二进制串形式），输出该数二进制表示中 1 的个数。例如，把 9表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。

思路：与1相与，可验证最后一位是否为1，然后再右移
*/
package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(9))
}

func hammingWeight(num uint32) int {
	var count int
	for num > 0 {
		count += int(num & 1)
		num = num >> 1
	}
	return count
}
