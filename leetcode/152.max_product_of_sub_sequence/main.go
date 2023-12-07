/*
152. 乘积最大子数组
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个 32-位 整数。

子数组 是数组的连续子序列。

示例 1:

输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。

思路：动态规划。但这里需要特殊注意，nums[i]的正负性是有影响的。
所以要用fmax[i]表示以数字i结尾的子数组的最大乘积。而fmin[i]表示最小乘积。
然后根据nums[i]，分正负性讨论。
则有当nums[i]为正数时，fmax[i]=max(fmax[i-1]*nums[i],nums[i])，fmin[i]=min(fmin[i-1]*nums[i],nums[i])
则有当nums[i]为负数时，fmax[i]=max(fmin[i-1]*nums[i],nums[i])，fmin[i]=min(fmax[i-1]*nums[i],nums[i])
继续递推即可
*/
package main

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	fmax, fmin := make([]int, len(nums)), make([]int, len(nums))
	fmax[0] = nums[0]
	fmin[0] = nums[0]
	res := fmax[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			fmax[i] = max(fmax[i-1]*nums[i], nums[i])
			fmin[i] = min(fmin[i-1]*nums[i], nums[i])
		} else if nums[i] < 0 {
			fmax[i] = max(fmin[i-1]*nums[i], nums[i])
			fmin[i] = min(fmax[i-1]*nums[i], nums[i])
		} else { //nums[i]==0
			fmax[i] = 0
			fmin[i] = 0
		}
		if fmax[i] > res {
			res = fmax[i]
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func main() {

}
