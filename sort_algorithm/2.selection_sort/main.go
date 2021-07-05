/**
选择排序（不稳定排序）
平均时间O(n^2), 最差O(n^2)，额外空间O(1)

原理概要：从i=0->size-1，不断循环，选出最小的放在对应的位置上

*/
package main

import "fmt"

func main() {
	data := []int{1, 5, 3, 6, 8, 12, 7, 4, 23, 9}
	selectionSort(data)
	fmt.Println(data)
}

func selectionSort(data []int) {
	size := len(data)
	var least int
	for i := 0; i < size; i++ {
		least = i
		for j := i + 1; j < size; j++ {
			if data[least] > data[j] {
				least = j
			}
		}
		data[i], data[least] = data[least], data[i]
	}
}
