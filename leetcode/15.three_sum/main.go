/*
*
15. 三数之和

给你一个整数数组 nums ，
判断是否存在三元组 [nums[i], nums[j], nums[k]]
满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。

思路：排序，然后三个指针，固定一个，剩下的两指针转化为2Sum问题
去重的思路：对于第一个和第二个指针，如果下一个寻找的数和当前数相等，就直接跳过下一个数，这样就可以去重了。
*/
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if nums[j]+nums[k] < -nums[i] {
				j++
			} else if nums[j]+nums[k] > -nums[i] {
				k--
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
