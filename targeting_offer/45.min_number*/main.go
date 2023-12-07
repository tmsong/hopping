/*
*
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

示例 1:

输入: [10,2]
输出: "102"
示例 2:

输入: [3,30,34,5,9]
输出: "3033459"

思路：此题求拼接起来的最小数字，本质上是一个排序问题。设数组 nums中任意两数字的字符串为 x 和 y ，则规定 排序判断规则 为：

若拼接字符串 x+y>y+x 则 x “大于” y ；
反之，若 x+y<y+x 则 x “小于” y ；
x “小于” y 代表：排序完成后，数组中 x 应在 y 左边；“大于” 则反之。

根据以上规则，套用任何排序方法对 nums 执行排序即可。
*/
package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i]*numLen(nums[j])+nums[j] < nums[j]*numLen(nums[i])+nums[i] {
			return true
		} else {
			return false
		}
	})
	fmt.Println(nums)
	buf := bytes.Buffer{}
	for _, num := range nums {
		buf.WriteString(strconv.Itoa(num))
	}
	return buf.String()
}

func numLen(num int) (digit int) {
	if num == 0 {
		return 10
	}
	for digit = 1; num > 0; num = num / 10 {
		digit = digit * 10
	}
	return digit
}

func main() {
	fmt.Println(minNumber([]int{1, 0}))
}
