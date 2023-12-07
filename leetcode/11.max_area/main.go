/*
11. 盛最多水的容器

给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。

第i和第j条线围成的面积：(j-i)*min(height[i],height[j])

两个指针一个指右一个指左，
如果左边高度小于右，则左指针往右移，如果右边高度小于左，则右指针往左移。
直到两个指针碰上。最大值将在这些过程中产出。


*/

package main

import (
	"fmt"
)

func maxArea(height []int) int {
	if len(height) < 2 { //特殊情况讨论
		return 0
	}
	var max int
	for i, j := 0, len(height)-1; i < j; {
		if res := (j - i) * min(height[i], height[j]); res > max {
			max = res
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
