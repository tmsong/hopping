/**
堆排序

原理概要：
用数组表示堆的特性：
1. 一个下标为n的节点，它的两个叶子结点的下标为2n+1 & 2n+2 （没越界的情况下）
2. 最后一个非叶子节点的下标 (len/2)-1

建堆：
从最后一个非叶子节点的下标开始，如果其值小于它的叶子结点，就将其与叶子结点交换。。。直到不能交换为止（moveDown）。
将堆顶与最后一个值交换，然后移动新堆顶，使其恢复成为一个有n-1个元素的堆。。。以此类推
*/
package main

import "fmt"

func main() {
	data := []int{1, 5, 3, 6, 8, 12, 7, 4, 23, 9}
	heapSort(data)
	fmt.Println(data)
}

func heapSort(data []int) {
	size := len(data)
	//建堆
	for i := size/2 - 1; i >= 0; i-- {
		moveDown(data, i, size-1)
	}
	//取大顶，和当前最后一位交换，然后恢复堆
	for i := size - 1; i >= 1; i-- {
		data[0], data[i] = data[i], data[0]
		moveDown(data, 0, i-1)
	}
}

func moveDown(data []int, first, last int) {
	largest := first*2 + 1 //找到first的左叶子结点
	for largest <= last {  //没越界就一直往下换
		if largest < last && data[largest] < data[largest+1] { //找到first的最大的叶子结点
			largest = largest + 1
		}
		if data[first] < data[largest] {
			data[first], data[largest] = data[largest], data[first] //交换first和largest
			first = largest                                         //以largest开始继续交换下去
			largest = 2*first + 1
		} else {
			break
		}
	}
}
