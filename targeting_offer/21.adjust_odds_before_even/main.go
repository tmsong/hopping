/**
剑指 Offer 21. 调整数组顺序使奇数位于偶数前面

输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。

思路：当然是左右指针，交换。

*/
package main

import "fmt"

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4, 5, 6, 7, 8}))
}

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]%2 == 1 { //奇数，符合预期，继续
			left++
			continue
		}
		if nums[right]%2 == 0 { //偶数，符合预期，继续
			right--
			continue
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	return nums
}
