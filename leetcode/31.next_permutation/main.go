/*
31.下一个排列

整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。
例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。

思路：下一个排列，经典问题，从右往左找到第一个比左边数字大的数字i，
然后再往右往左找到第一个比这个数字大的数字j，交换这两个数字i,j，
然后反向所有原本i位置后面的数字（不包括i）。
*/

package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	i := len(nums) - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- { //从右往左找到第一个比右边数字小的数字
	}
	if i < 0 { //说明已经是最大排列，直接重置到最小排列
		reverse(nums)
		return
	}
	j := len(nums) - 1
	for ; nums[j] <= nums[i]; j-- { //从右往左找到第一个比这个数字大的数字
	}
	nums[i], nums[j] = nums[j], nums[i] //交换这两个数字
	reverse(nums[i+1:])                 //再反向所有i位置后的数字（不包括i）
	return
}

func reverse(nums []int) {
	if len(nums) < 2 {
		return
	}
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	nextPermutation(nums)
	fmt.Println(nums)
}
