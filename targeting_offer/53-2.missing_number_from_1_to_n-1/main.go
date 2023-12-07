/*
*
剑指 Offer 53 - II. 0～n-1中缺失的数字

一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。

在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

示例 1:

输入: [0,1,3]
输出: 2

示例 2:

输入: [0,1,2,3,4,5,6,7,9]
输出: 8

思路：还是二分法查找，因为这个数字之后的数字，都有nums[i] = i+1
*/
package main

import "fmt"

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right-1 {
		mid := (left + right) / 2
		if nums[mid] == mid {
			left = mid
		} else {
			right = mid
		}
	}
	if nums[left] == left && nums[right] == right+1 { //证明这个数已经被逼在left和right指针的中间，那么这个数应该等于right
		return right
	}
	//再讨论两个数的极性相等的时候
	if nums[right] == right { //这时，数在right右边一个
		return right + 1
	} else if nums[left] == left+1 { //这时，数就是left
		return left
	}
	return 0
}

func main() {
	fmt.Println(missingNumber([]int{1, 2, 3}))
}
