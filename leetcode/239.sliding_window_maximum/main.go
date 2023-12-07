/*
*
239. 滑动窗口最大值

科技馆内有一台虚拟观景望远镜，它可以用来观测特定纬度地区的地形情况。
该纬度的海拔数据记于数组 heights ，其中 heights[i] 表示对应位置的海拔高度。
请找出并返回望远镜视野范围 limit 内，可以观测到的最高海拔值。

示例 1：

输入：heights = [14,2,27,-5,28,13,39], limit = 3
输出：[27,27,28,28,39]
解释：

	滑动窗口的位置                最大值

---------------               -----
[14 2 27] -5 28 13 39          27
14 [2 27 -5] 28 13 39          27
14 2 [27 -5 28] 13 39          28
14 2 27 [-5 28 13] 39          28
14 2 27 -5 [28 13 39]          39

思路1：
用一个数据结构保存每个数以及这个数对应的下标，
创建一个以元素为基础的大顶堆，
随着视野移动，不断往里push元素，但是，一旦发现此时堆的top的数对应的下标超过视野，则弹出这个顶元素。

思路2：
把大顶堆换成队列，保存的是下标。目标是让队列中的下标对应的元素单调递减（队头大，队尾小）
这样，每次往队尾添加下标时，如果对应元素大于前面下标对应的元素，就可以移除前面的。（比思路1中大顶堆需要维护的元素数量少）
同样，一旦发现此时队头的数对应的下标超过视野，则弹出这个队头元素。
*/
package main

import "fmt"

type Telescope struct {
	Heights []int
	Queue   []int
	Limit   int
}

func (t *Telescope) Push(index int) {
	if len(t.Queue) == 0 {
		t.Queue = append(t.Queue, index)
		return
	}
	for i := len(t.Queue) - 1; i >= 0; i-- {
		if t.Heights[t.Queue[i]] < t.Heights[index] { //每次往队尾添加元素时，如果这个元素大于前面已添加的元素，就可以移除前面的元素。
			t.Queue = t.Queue[:i]
		} else {
			break
		}
	}
	t.Queue = append(t.Queue, index)
	for len(t.Queue) > 0 && t.Queue[0]+t.Limit <= index { //一旦发现此时队头的数对应的下标超过视野，则弹出这个队头元素。
		t.Queue = t.Queue[1:]
	}
}

func (t *Telescope) Top() int {
	if len(t.Queue) == 0 {
		return -1
	}
	return t.Heights[t.Queue[0]]
}

func maxAltitude(heights []int, limit int) []int {
	if limit == 0 {
		return nil
	}
	ts := &Telescope{Heights: heights, Limit: limit}
	for i := 0; i < limit; i++ {
		ts.Push(i)
	}
	ret := []int{ts.Top()}
	for i := limit; i < len(heights); i++ {
		ts.Push(i)
		ret = append(ret, ts.Top())
	}
	return ret
}

func main() {
	fmt.Println(maxAltitude([]int{14, 2, 27, -5, 28, 13, 39}, 3))
}
