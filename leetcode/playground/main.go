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

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	c1, uc1 := robWithNoCondition(nums[:len(nums)-1])
	res1 := max(c1, uc1+nums[len(nums)-1])
	c2, uc2 := robWithNoCondition(nums[1:])
	res2 := max(c2, uc2+nums[0])
	return max(res1, res2)
}

func robWithNoCondition(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	} else if len(nums) == 1 {
		return nums[0], 0
	}
	//一个用于记录选i时可取得的最大值，则有chosen[i] = max(notChosen[i-1]+nums[i], chosen[i-i])
	chosen := make([]int, len(nums))
	//一个用于记录没选i时可取得的最大值，则有notChosen[i] = max(chosen[i-1], notChosen[i-i])
	notChosen := make([]int, len(nums)) //一个用于记录选了i时可取得的最大值
	chosen[0] = nums[0]
	notChosen[0] = 0
	for i := 1; i < len(chosen); i++ { //先取第一间，然后算到第n-1
		chosen[i] = max(notChosen[i-1]+nums[i], chosen[i-i])
		notChosen[i] = max(chosen[i-1], notChosen[i-i])
	}
	return chosen[len(chosen)-1], notChosen[len(notChosen)-1] //最后返回两种情况可能达到的最大值
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
}
