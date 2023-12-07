/*
*
213. 打家劫舍 II
已解答
中等
相关标签
相关企业
提示
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

示例 1：

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2：

输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。

	偷窃到的最高金额 = 1 + 3 = 4 。

思路：
打家劫舍I，上一间选了这一间就不能选，所以有递推公式：
dp[i] = max(dp[i-2]+nums[i],dp[i-1])
和打家劫舍I的区别就是，如果第一间房选了，则最后一间房不能选，如果第一间房没选，则可以选最后一间房。
所以可以计算两次，一次计算从0->n-2，一次计算从1->n-1，取更大值。
*/
package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	return max(robWithNoCondition(nums[1:]), robWithNoCondition(nums[:len(nums)-1]))
}

func robWithNoCondition(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[0] = nums[0]
		} else if i == 1 {
			dp[1] = max(nums[0], nums[1])
		} else {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}
	}
	return dp[len(nums)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
