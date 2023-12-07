/*
56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

思路：按照区间的第一个点从小到大排序，然后按照顺序一个一个往这个上merge，
如果和前一个merge不了，则一定是一个新的区间。证明略
*/
package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ret := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if len(ret) == 0 {
			ret = append(ret, intervals[i])
			continue
		}
		l := len(ret)
		if intervals[i][0] <= ret[l-1][1] { //说明可合并
			ret[l-1][1] = max(intervals[i][1], ret[l-1][1])
		} else {
			ret = append(ret, intervals[i])
		}
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	fmt.Println()
}
