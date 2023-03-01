/*
剑指 Offer 44. 数字序列中某一位的数字

数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

请写一个函数，求任意第n位对应的数字。

思路：先确定落到几位数，再确定是第几个数的第几位
10个一位数，100-10个两位数，1000-100个三位数...
*/
package main

import "fmt"

func findNthDigit(n int) int {
	if n == 0 { //特殊处理第一位的0
		return 0
	}
	digits, p := 1, 1    //digits代表当前落到几位数的范围内了，p表示10的幂
	for n > digits*p*9 { //表示当前落在了digits位数中了，跳出
		n -= digits * p * 9 //否则减掉前面所有的digits位数，进入digits+1位数中寻找
		p = p * 10
		digits++
	}
	//计算是p后的第几个数字
	if i := n % digits; i == 0 {
		return (p + n/digits - 1) % 10 //说明就是第p+n/digit-1个数的最后一位
	} else {
		x := p + n/digits //否则，是 p + n/digits 的第 n%digits 位
		for digits-i > 0 {
			x = x / 10
			i++
		}
		return x % 10
	}
}

func main() {
	fmt.Println(findNthDigit(3))
}
