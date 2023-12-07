/*
*
剑指 Offer 56 - I. 数组中数字出现的次数

一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

示例 1：

输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]

示例 2：

输入：nums = [1,2,10,4,1,4,3,3]
输出：[2,10] 或 [10,2]

思路：
如果只有一个数字只出现了一次，那么可以用二进制异或的方式，这样其他所有出现了两次的数字都会抵消，只剩下这个出现一次的数字。
现在有两个数字只出现了一次，那么就设法把所有数字分成两组，使得：
1.两个只出现一次的数字在不同的组中；
2.相同的数字会被分到相同的组中。

办法：
1.将所有数字先异或一次，得到这两个数的异或结果。
2.在异或结果（二进制）里任取一位=1的位（表示a与b在这一位上不同）。
3.用这个位来区分所有数（位与结果为0或1），将数分成两组。
4.每组分别异或，就能得到a和b。
*/
package main

import "fmt"

func singleNumbers(nums []int) []int {
	var allOr int
	for _, num := range nums {
		allOr = allOr ^ num //先把所有数都异或一遍
	}
	digit := 1
	if allOr == 0 { //没找到，这个数组有问题。
		return nil
	}
	for allOr%2 == 0 {
		digit *= 2 //从右到左找到异或结果里第一个为1的位
		allOr /= 2
	}
	//开始异或两个
	var a, b int
	for _, num := range nums {
		if num&digit == 0 { //那一位为0
			a = a ^ num
		} else {
			b = b ^ num
		}
	}
	return []int{a, b}
}

func main() {
	fmt.Println(singleNumbers([]int{1, 2, 10, 4, 1, 4, 3, 3}))
}
