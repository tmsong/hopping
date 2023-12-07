/*
*
剑指 Offer 53 - I. 在排序数组中查找数字 I

统计一个数字在排序数组中出现的次数。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2

示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0

思路：排序数组，二分法查找
*/
package main

import "fmt"

func search(nums []int, target int) int {
	l := len(nums)
	if l <= 0 { //长度为0，直接返回
		return 0
	} else if l == 1 { //长度为1，分情况讨论
		if nums[0] == target {
			return 1
		} else {
			return 0
		}
	}
	mid := nums[l/2]
	if mid > target {
		return search(nums[:l/2], target) //在左边
	} else if mid < target {
		return search(nums[l/2+1:], target) //在右边
	} else {
		var count int
		for ptr := l / 2; ptr >= 0; ptr-- { //就在中间，分别往左和往右拓展计数
			if nums[ptr] == target {
				count++
			} else {
				break
			}
		}
		for ptr := l/2 + 1; ptr < len(nums); ptr++ {
			if nums[ptr] == target {
				count++
			} else {
				break
			}
		}
		return count
	}
}

func main() {
	fmt.Println(search([]int{2, 2}, 2))
}
