/*
*
剑指 Offer 56 - II. 数组中数字出现的次数 II

在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

示例 1：

输入：nums = [3,4,3,3]
输出：4
示例 2：

输入：nums = [9,1,7,9,7,9,7]

思路：统计所有数字的各二进制位中 1 的出现次数，并对 3 求余，结果则为只出现一次的数字。

方法：有限状态自动机，可以不用if来判断每一个数的每一位，那样复杂度就比较高。
对于每一位，都有三种状态：0，1，2，当超过2时，又将回到0。
三种状态，就需要用两位来记录状态的转换，也就是说，需要两个数，这里用one，two表示。
也就是，用one的第n位值拼上two的第n位值，等于原数组中第n位中1出现的次数对3求余的结果：

two one
0   0    = 0次，当新数对应位为1，则变为0，1
0   1    = 1次，当新数对应位为1，则变为1，0
1   0    = 2次，当新数对应位为1，则变为0，0

则不难得出，当处理一个新数x时，每一位
one = (one^x) & (!two)
two = (two^x) & (two|one)
*/
package main

import "fmt"

func singleNumber(nums []int) int {
	var one, two int
	for _, num := range nums {
		one, two = (one^num)&(^two), (two^num)&(one|two)
	}
	return one
}

func main() {
	fmt.Println(singleNumber([]int{12, 1, 6, 12, 6, 12, 6}))
}
