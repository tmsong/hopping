/*
*
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

示例 1:

输入: [7,5,6,4]
输出: 5

思路1：（复杂度太高被否了）插入排序（从大到小），对于每个数而言，在第几个插入数组就表示前面有几个比他大的数字。

思路2： 归并排序，两个排好序的数组在合并时，
L = [8, 12, 16, 22, 100]   R = [9, 26, 55, 64, 91]  M = []

	  |                          |
	lPtr                       rPtr

我们发现 lPtr 指向的元素小于 rPtr 指向的元素，于是把 lPtr 指向的元素放入答案，并把 lPtr 后移一位。

L = [8, 12, 16, 22, 100]   R = [9, 26, 55, 64, 91]  M = [8]

	  |                       |
	lPtr                     rPtr

这个时候我们把左边的 888 加入了答案，我们发现右边没有数比 8 小，所以 8 对逆序对总数的「贡献」为 0。

接着我们继续合并，把 999 加入了答案，此时 lPtr 指向 12，rPtr 指向 26。

L = [8, 12, 16, 22, 100]   R = [9, 26, 55, 64, 91]  M = [8, 9]

	 |                          |
	lPtr                       rPtr

此时 lPtr 比 rPtr 小，把 lPtr 对应的数加入答案，并考虑它对逆序对总数的贡献为 rPtr 相对 R 首位置的偏移 1（即右边只有一个数比 12 小，所以只有它和 12 构成逆序对），以此类推。

我们发现用这种「算贡献」的思想在合并的过程中计算逆序对的数量的时候，只在 lPtr 右移的时候计算，是基于这样的事实：当前 lPtr 指向的数字比 rPtr 小，
但是比 R 中 [0 ... rPtr - 1] 的其他数字大，[0 ... rPtr - 1] 的其他数字本应当排在 lPtr 对应数字的左边，但是它排在了右边，所以这里就贡献了 rPtr 个逆序对。
*/
package main

import "fmt"

func mergeSort(nums, tmp []int, l, r int) (reverseCnt int) { //左闭右开
	if r-l <= 1 {
		return 0
	}
	mid := (l + r) / 2
	reverseCnt = mergeSort(nums, tmp, l, mid) + mergeSort(nums, tmp, mid, r)
	lptr, rptr, tmpPos := l, mid, l
	for lptr < mid && rptr < r {
		if nums[lptr] > nums[rptr] {
			tmp[tmpPos] = nums[rptr]
			rptr++
		} else {
			reverseCnt += (rptr - mid) //左指针的元素小于等于右指针的元素，却一定大于右指针左边的元素，所以右指针左边的元素每一个都跟左指针的元素构成了逆序对
			tmp[tmpPos] = nums[lptr]
			lptr++
		}
		tmpPos++
	}
	for lptr < mid {
		tmp[tmpPos] = nums[lptr]
		reverseCnt += (rptr - mid) //同理
		tmpPos++
		lptr++
	}
	for rptr < r {
		tmp[tmpPos] = nums[rptr]
		tmpPos++
		rptr++
	}
	copy(nums[l:r], tmp[l:r])
	return reverseCnt
}

func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	tmp := make([]int, len(nums))
	return mergeSort(nums, tmp, 0, len(nums))
}

func main() {
	nums := []int{1, 3, 2, 3, 1}
	fmt.Println(reversePairs(nums))
}
