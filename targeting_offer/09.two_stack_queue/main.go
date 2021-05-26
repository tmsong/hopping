/**
剑指 Offer 09 用两个栈实现队列

用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead操作返回 -1 )

思路：
这个题好像没什么意义，就来实现个循环数组吧。
*/
package main

import "fmt"

func main() {
	queue := Constructor()
	for i := 0; i < 10000; i++ {
		queue.AppendTail(i)
	}
	for i := 0; i < 50; i++ {
		fmt.Println(queue.DeleteHead())
	}
	for i := 0; i < 100; i++ {
		queue.AppendTail(i)
	}
	fmt.Println(queue.Len)
	fmt.Println(queue.Head)
}

type CQueue struct {
	Head   int
	Len    int
	Values []int
}

func Constructor() CQueue {
	return CQueue{
		Values: make([]int, 10000),
	}
}

func (this *CQueue) AppendTail(value int) {
	if this.Len >= 10000 {
		return
	}
	this.Values[(this.Head+this.Len)%10000] = value
	this.Len++
}

func (this *CQueue) DeleteHead() int {
	if this.Len <= 0 {
		return -1
	}
	defer func() {
		this.Head = (this.Head + 1) % 10000
		this.Len--
	}()
	return this.Values[this.Head]
}
