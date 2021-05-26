/**
剑指 Offer 03. 数组中重复的数字

找出数组中重复的数字。

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3

思路：
如果没有重复数字，那么正常排序后，数字i应该在下标为i的位置，所以思路是重头扫描数组，遇到下标为i的数字如果不是i的话，（假设为m)，那么我们就拿与下标m的数字交换。
在交换过程中，如果有重复的数字发生，那么终止返回。
*/
package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}

func findRepeatNumber(nums []int) int {
	for idx := range nums {
		if nums[idx] != idx {
			if nums[idx] == nums[nums[idx]] {
				return nums[idx]
			}
			nums[idx], nums[nums[idx]] = nums[nums[idx]], nums[idx]
		}
	}
	return -1
}
