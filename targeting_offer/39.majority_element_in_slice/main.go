/**
剑指 Offer 39. 数组中出现次数超过一半的数字

数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例1:

输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2

思路：最笨的方法当然是直接统计数字出现的个数，
但实际上，其实可以用打擂台的方式，也就是记当前出现最多次数的数字，出现一次+1
如果出现其他的数字，就把这个数字的值打掉1。
*/
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	var curNum, curTime int
	for _, num := range nums {
		if curTime == 0 {
			curNum, curTime = num, 1
		} else if curNum == num {
			curTime++
		} else if curNum != num {
			curTime--
		}
	}
	return curNum
}
