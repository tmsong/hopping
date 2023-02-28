/*
剑指 Offer 43. 1～n 整数中 1 出现的次数

输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

思路：从个位开始逐位计算
设整数为x，以个位为第0位，十位为第1位，依次递推，不难得出第n位出现1的次数为：
如果(x/10^n)%10，即第n位的数字，大于1时，第n位出现1的次数为：(x/10^(n+1))*10^n +  10^n
如果(x/10^n)%10，即第n位的数字，等于1时，第n位出现1的次数为：(x/10^(n+1))*10^n + (x%(10^n)+1)
如果(x/10^n)%10，即第n位的数字，小于1时，第n位出现1的次数为：(x/10^(n+1))*10^n
*/
package main

import "fmt"

func countDigitOne(x int) int {
	//tmp = 10^n
	tmp, count := 1, 0
	for {
		if x/tmp == 0 {
			break
		}
		count += (x / (tmp * 10)) * tmp
		if digit := (x / tmp) % 10; digit > 1 {
			count += tmp
		} else if digit == 1 {
			count += (x % tmp) + 1
		}
		tmp = tmp * 10
	}
	return count
}

func main() {
	fmt.Println(countDigitOne(88))
}
