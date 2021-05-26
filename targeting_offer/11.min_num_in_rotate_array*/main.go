/**
剑指 Offer 11. 旋转数组的最小数字

把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

示例 1：

输入：[3,4,5,1,2]
输出：1
示例 2：

输入：[2,2,2,0,1]
输出：0

思路：
首尾指针，二分法。始终保持左针>右针，则最小元素在左右针之间。
*/
package main

import "fmt"

func main() {
	fmt.Println(minArray([]int{1, 3, 5}))
}

func minArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	for i, j := 0, len(numbers)-1; ; {
		if i == j {
			return numbers[i]
		}
		middle := (i + j) / 2
		if numbers[middle] > numbers[j] {
			i = middle + 1
		} else if numbers[middle] < numbers[j] {
			j = middle
		} else {
			j--
		}
	}
	return numbers[0]
}
