/*
剑指 Offer 42. 连续子数组的最大和

输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。

示例1:

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释:连续子数组[4,-1,2,1] 的和最大，为6。

思路：时间复杂度要求低，那就用空间换时间，多存点东西解决。
多一个maxSum数组，其中maxSum[i]保存以i结尾的数组中的的最大和。
转移方程：maxSum[i] = maxSum[i-1]+nums[i] 或 nums[i]
*/
package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, lastMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if lastMax > 0 {
			lastMax = lastMax + nums[i]
		} else {
			lastMax = nums[i]
		}
		if lastMax > max {
			max = lastMax
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
