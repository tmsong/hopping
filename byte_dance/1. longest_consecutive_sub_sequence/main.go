/**
标题：最长连续子序列
描述信息
题目：

给定一个整形数组长度为n

求构造出子序列连续并且长度最大，要求子序列有顺序。

连续定义序列 [x, x+1, ..., x+k-1] ，长度是k

例如数组 [5, 3, 1, 2, 4] 的子序列 正确的表述例子， [3] [5,3, 1, 2, 4] [5, 1, 4] 但是 [1, 3] 不是，因为要按照从左往右顺序 [3, 1] 是正确的子序列。

输入：

n个整数，1<= n <= 200000

a1, a2, ..., an (1<= ai <= 10^9 ）

输出：

k

任意一个可以组成最大长度的连续子序列的数组元素下标 数组，下标按照递增排列

Example：

input

7

3 3 4 7 5 6 8

output

4

2 3 5 6

-----------

input

16

6 7 8 3 4 5 9 10 11 6 7 8 9 6 10 11

output

9

4 5 6 10 11 12 13 15 16

思路：
复杂度 O(n)
设 zip[input[i] , i] 为以 数组元素为key， 下标为value

设以i 下标为结尾的数组 最大子序列 cnt[i] ,

如果[0, ..., i+1] 序列里存在 inputs[i+1]-1 ，

那么 cnt[i+1] = cnt[zip[inputs[i+1]-1]] + 1 相等于把前面的片段接起来

利用map记前面元素的idx，遍历一遍就可以解决

*/
package main

import (
	"fmt"
)

func main() {
	seq := []int64{6, 7, 8, 3, 4, 5, 9, 10, 11, 6, 7, 8, 9, 6, 10, 11}
	size, idxs := longestConsecutiveSubSequence(seq)
	fmt.Println(size)
	fmt.Println(idxs)
}

func longestConsecutiveSubSequence(seq []int64) (size int, idxs []int) {
	//元素:->元素出现过的下标集合
	valueIdxsMap := make(map[int64]int)
	//从对应下标元素往前数，能连成连续序列的元素个数
	maxCnts := make([]int, len(seq))
	//从对应下标元素往前数，上一个与之能连成连续序列的元素的下标
	lastIdxs := make([]int, len(seq))
	for idx, value := range seq {
		valueIdxsMap[value] = idx
		lastIdxs[idx] = idx
		maxCnts[idx] = 1                              //默认的序列长度就是1
		if prevIdx, ok := valueIdxsMap[value-1]; ok { //在map里找序列前面有没有出现过value-1的元素
			maxCnts[idx] = maxCnts[prevIdx] + 1
			lastIdxs[idx] = prevIdx
		}
	}
	//循环完毕，找到最大的cnt
	var lastIdx int
	for idx, cnt := range maxCnts {
		if cnt > size {
			size = cnt
			lastIdx = idx
		}
	}
	idxs = make([]int, size)
	i := size - 1
	for i >= 0 {
		idxs[i] = lastIdx
		i--
		lastIdx = lastIdxs[lastIdx]
	}
	return
}
