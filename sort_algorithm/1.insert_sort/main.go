/**
插入排序（稳定排序）
平均时间O(n^2), 最差O(n^2)，额外空间O(1)

原理概要：
很简单，固定i从0->len，取j=i，将data[j]插入到[0,i]间应处于的位置
*/
package main

import "fmt"

func main() {
	data := []int{1, 5, 3, 6, 8, 12, 7, 4, 23, 9}
	insertSort(data)
	fmt.Println(data)
}

func insertSort(data []int) {
	size := len(data)
	for i := 0; i < size; i++ {
		tmp := data[i]
		var j int
		for j = i; j > 0 && tmp < data[j-1]; j-- {
			data[j] = data[j-1] //把后面的值都往后挪一位，好执行插入
		}
		data[j] = tmp
	}
}
