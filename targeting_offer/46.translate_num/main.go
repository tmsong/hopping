/*
*
剑指 Offer 46. 把数字翻译成字符串

给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

思路：递归，按2或按1切
*/
package main

import "fmt"

func translateNum(num int) int {
	//如果小于10，则直接结束
	if num/10 == 0 {
		return 1
	}
	//大于10的情况下，看是两位数还是1位数。
	if num%100 < 10 { //后两位是0x的情况下，只能用1位
		return translateNum(num / 10)
	} else if num%100 < 26 { //可以是2位数，可以是1位数
		return translateNum(num/10) + translateNum(num/100)
	} else { //只能是1位数
		return translateNum(num / 10)
	}
}

func main() {
	fmt.Println(translateNum(506))
}
