/*
*
4. 寻找两个正序数组的中位数

给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
*/
package main

import "fmt"

/*
思路：先知道要找第几位数字，然后再二分法查找
根据中位数的定义，当 m+n 是奇数时，中位数是两个有序数组中的第(m+n)/2 个元素，
当 m+n 是偶数时，中位数是两个有序数组中的第 (m+n)/2 个元素和第 (m+n)/2+1 个元素的平均值。
因此，这道题可以转化成寻找两个有序数组中的第 k 小的数，其中 k 为 (m+n)/2或 (m+n)/2+1。

那么可以比较nums1[k/2-1]与nums2[k/2-1]，对于这两个数中的更小值（假设为nums1[k/2-1])，
最多只会有 (k/2−1)+(k/2−1)≤k−2 个元素比它小，那么它就不能是第k小的数了。可以刨除nums1中所有小于nums1[k/2-1]的数

于是，问题就转化为了求nums1[k/2: ]和nums2中第k/2小的元素
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//目标，以及总量是否是偶数。target = (l1+l2)/2因为总量是奇数时，是找第target+1个数，总量是偶数时，是找target和target+1
	target, isTwo := (len(nums1)+len(nums2))/2, (len(nums1)+len(nums2))%2 == 0
	if !isTwo {
		target += 1
	}
	for target > 1 && len(nums1) > 0 && len(nums2) > 0 {
		step := target/2 - 1
		if step >= len(nums1) { //处理step直接大于了一个数组的情况
			step = len(nums1) - 1
		} else if step >= len(nums2) {
			step = len(nums2) - 1
		}
		if nums1[step] <= nums2[step] { //这下nums1的前step+1个数都可以抛弃掉了
			nums1 = nums1[step+1:]
		} else {
			nums2 = nums2[step+1:] //反之亦然
		}
		target -= (step + 1) //问题转化为找nums1和nums2的第target-step-1个数
	}
	if len(nums1) == 0 || len(nums2) == 0 {
		var nums []int
		if len(nums1) == 0 {
			nums = nums2
		} else {
			nums = nums1
		}
		if isTwo { //总数为偶数的时候，需要中间两个数相加/2
			return float64(nums[target-1]+nums[target]) / 2
		} else {
			return float64(nums[target-1])
		}
	}
	//target = 1，且这两个数组都有值的情况
	if isTwo {
		var a, b int
		if nums1[0] > nums2[0] {
			a = nums2[0]
			nums2 = nums2[1:]
		} else {
			a = nums1[0]
			nums1 = nums1[1:]
		}
		if len(nums1) == 0 {
			b = nums2[0]
		} else if len(nums2) == 0 {
			b = nums1[0]
		} else if nums1[0] < nums2[0] {
			b = nums1[0]
		} else {
			b = nums2[0]
		}
		return float64(a+b) / 2
	} else {
		if nums1[0] < nums2[0] {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
