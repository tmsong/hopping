/*
归并排序（稳定排序）
平均时间O(nlogn), 最差O(n^2)，额外空间O(1)

原理概要：
1. 将序列二分，分别排好序后，合并两个已排好序的序列。
*/
package main

import "fmt"

func main() {
	data := []int{1, 6, 5, 3, 6, 8, 23, 7, 8, 12, 7, 5, 1, 3, 4, 23, 9}
	fmt.Println(mergeSort(data))
}

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	} else {
		l := len(data)
		return merge(mergeSort(data[:l/2]), mergeSort(data[l/2:]))
	}
}

func merge(left []int, right []int) []int {
	merged := make([]int, len(left)+len(right))
	var cur, i, j int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged[cur] = left[i]
			i++
		} else {
			merged[cur] = right[j]
			j++
		}
		cur++
	}
	for i < len(left) {
		merged[cur] = left[i]
		i++
		cur++
	}
	for j < len(right) {
		merged[cur] = right[j]
		j++
		cur++
	}
	return merged
}
