/*
*
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。

思路：维护一个大顶堆，一个小顶堆，各保存列表的一半元素，大顶堆存较小的那一半，小顶堆存较大的那一半。
进来元素的时候，轮流往大顶堆或小顶堆里推，保证两个堆里面的元素相等/或多一个。
当元素总体位奇数个的时候，往大顶堆推（或小顶堆），这样奇数的时候取大顶堆顶（或小顶堆顶），偶数的时候取两个顶求平均。
当往大顶堆推时，需要先推到小顶堆，然后弹出最小的元素（小顶堆顶），再往大顶堆推，这样能保证大顶堆里永远是较小的那一半，反之亦然。
*/
package main

import "fmt"

type MedianFinder struct {
	Count         int
	BigTopHeap    Heap
	LittleTopHeap Heap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		BigTopHeap:    Heap{Bigtop: true},
		LittleTopHeap: Heap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.Count%2 == 0 {
		this.LittleTopHeap.Push(num)
		this.BigTopHeap.Push(this.LittleTopHeap.Pop())
	} else {
		this.BigTopHeap.Push(num)
		this.LittleTopHeap.Push(this.BigTopHeap.Pop())
	}
	this.Count++
	return
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Count%2 == 0 {
		return (float64(this.BigTopHeap.Top()) + float64(this.LittleTopHeap.Top())) / 2
	} else {
		return float64(this.BigTopHeap.Top())
	}
}

type Heap struct {
	Bigtop   bool
	Elements []int
}

func (h *Heap) Push(num int) {
	h.Elements = append(h.Elements, num) //推元素到堆角
	pos := len(h.Elements) - 1
	for pos > 0 {
		up := (pos+1)/2 - 1                                                                                    //找到上面的元素
		if (h.Bigtop && h.Elements[up] < h.Elements[pos]) || (!h.Bigtop && h.Elements[up] > h.Elements[pos]) { //如果大了就交换
			h.Elements[up], h.Elements[pos] = h.Elements[pos], h.Elements[up]
		}
		pos = up
	}
}

func (h *Heap) Pop() int {
	if len(h.Elements) == 0 {
		return 0
	}
	ret := h.Elements[0]
	h.Elements[0] = h.Elements[len(h.Elements)-1]
	h.Elements = h.Elements[:len(h.Elements)-1]
	//接下来，恢复堆
	first, largest, last := 0, 1, len(h.Elements)-1
	for largest <= last {
		if largest < last && ((h.Bigtop && h.Elements[largest] < h.Elements[largest+1]) || (!h.Bigtop && h.Elements[largest] > h.Elements[largest+1])) {
			largest = largest + 1
		} //往下交换
		if (h.Bigtop && h.Elements[first] < h.Elements[largest]) || (!h.Bigtop && h.Elements[first] > h.Elements[largest]) {
			h.Elements[first], h.Elements[largest] = h.Elements[largest], h.Elements[first]
			first = largest
			largest = 2*first + 1
		} else {
			break
		}
	}
	return ret
}

func (h *Heap) Top() int {
	if len(h.Elements) == 0 {
		return 0
	}
	return h.Elements[0]
}

func main() {
	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	fmt.Println(m.FindMedian())
	m.AddNum(3)
	fmt.Println(m.FindMedian())
	m.AddNum(5)
	fmt.Println(m.FindMedian())
}
