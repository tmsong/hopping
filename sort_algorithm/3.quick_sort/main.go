/*
*
快速排序（不稳定排序）
平均时间O(nlog2n), 最差O(n^2)，额外空间O(1)

原理概要：
1. 设定一个分界值key，通过移动左右指针，找到左指针大于key的元素，右指针指向小于key的元素，交换，直至左指针=右指针。
2. 不断递归重复1的操作。
*/
package main

import "fmt"

func main() {
	data := []int{1, 6, 5, 3, 6, 8, 23, 7, 8, 12, 7, 5, 1, 3, 4, 23, 9}
	quickSort(data)
	fmt.Println(data)
}

func quickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	key := data[0]
	i, j := 1, len(data)-1
	for i < j {
		for data[i] <= key && i < len(data)-1 {
			i++
		}
		for data[j] >= key && j > 0 {
			j--
		}
		if j > i {
			data[i], data[j] = data[j], data[i]
		}
	}
	data[0], data[i] = data[i], data[0] //交换key和左边最大的值
	quickSort(data[:i])
	quickSort(data[j:])
}
