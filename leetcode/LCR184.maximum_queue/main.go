/*
请设计一个自助结账系统，该系统需要通过一个队列来模拟顾客通过购物车的结算过程，需要实现的功能有：

get_max()：获取结算商品中的最高价格，如果队列为空，则返回 -1
add(value)：将价格为 value 的商品加入待结算商品队列的尾部
remove()：移除第一个待结算的商品价格，如果队列为空，则返回 -1
注意，为保证该系统运转高效性，以上函数的均摊时间复杂度均为 O(1)

示例 1：

输入:
["Checkout","add","add","get_max","remove","get_max"]
[[],[4],[7],[],[],[]]

输出: [null,null,null,7,4,7]
示例 2：

输入:
["Checkout","remove","get_max"]
[[],[],[]]

输出: [null,-1,-1]

思路：维护两个队列，
一个储存元素A，
一个单调队列B，按元素从大到小，但储存的是对应的下标。
插入元素：A队尾加入，B队尾去掉所有比这个元素小的元素下标后，再添加此元素下标。
弹出元素：A队首弹出，然后检查B的队首，弹出下标在这个元素前的元素，直到遇到一个下标在后面的。
*/
package main

type Checkout struct {
	Queue         []int
	PriorityIndex []int
	MinIndex      int //记录当前的最小下标（此值等于已弹出了多少个元素）
}

func Constructor() Checkout {
	return Checkout{
		Queue:         []int{},
		PriorityIndex: []int{},
	}
}

func (this *Checkout) Get_max() int {
	if len(this.Queue) == 0 {
		return -1
	}
	return this.Queue[this.PriorityIndex[0]-this.MinIndex]
}

func (this *Checkout) Add(value int) {
	index := len(this.Queue) + this.MinIndex //当前元素的真实下标
	this.Queue = append(this.Queue, value)
	if len(this.PriorityIndex) > 0 {
		i := len(this.PriorityIndex) - 1 //从单调队列的队尾去掉所有比这个元素小的元素下标。
		for ; i >= 0 && this.Queue[this.PriorityIndex[i]-this.MinIndex] < value; i-- {
		}
		this.PriorityIndex = this.PriorityIndex[:i+1]
	}
	this.PriorityIndex = append(this.PriorityIndex, index)
}

func (this *Checkout) Remove() int {
	if len(this.Queue) == 0 {
		return -1
	}
	ret := this.Queue[0]
	this.Queue = this.Queue[1:]
	this.MinIndex++
	if len(this.PriorityIndex) > 0 {
		i := 0
		for ; i < len(this.PriorityIndex) && this.PriorityIndex[i] < this.MinIndex; i++ {
		}
		this.PriorityIndex = this.PriorityIndex[i:] //检查B的队首，弹出所有下标在这个元素前的元素，直到遇到一个下标在后面的。
	}
	return ret
}

func main() {
	x := []int{0}
	x = x[1:]
	x = x[1:]
}
