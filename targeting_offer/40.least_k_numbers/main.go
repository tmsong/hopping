/**
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]

思路1：直接通过快排切分排好第 K 小的数（下标为 K-1），那么它左边的数就是比它小的另外 K-1 个数。
第一次切分的时候需要遍历整个数组 (0 ~ n) 找到了下标是 j 的元素，
假如 k 比 j 小的话，那么我们下次切分只要遍历数组 (0~k-1)的元素就行啦，
反之如果 k 比 j 大的话，那下次切分只要遍历数组 (k+1～n) 的元素就行。（有点类似于快排+二分法）
思路2：维护一个K个数的大顶堆，如果当前堆元素不足K，则直接入堆，否则跟堆顶比，比堆顶小就替换掉堆顶，然后恢复堆。
*/
package main

import "fmt"

func main() {
	arr := []int{3, 2, 1, 4, 5, 0}
	fmt.Println(getLeastNumbers(arr, 2))
}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	//建堆
	for i := k/2 - 1; i >= 0; i-- {
		moveDown(arr, i, k)
	}
	//后面的元素过一遍
	for i := k; i < len(arr); i++ {
		if arr[i] < arr[0] {
			arr[i], arr[0] = arr[0], arr[i]
			moveDown(arr, 0, k)
		}
	}
	return arr[:k]
}

func moveDown(arr []int, cur, last int) {
	child := cur*2 + 1
	for child < last {
		if child+1 < last && arr[child] < arr[child+1] {
			child++
		}
		if arr[cur] < arr[child] {
			arr[cur], arr[child] = arr[child], arr[cur]
			cur = child
			child = cur*2 + 1
		} else {
			break
		}
	}
}
