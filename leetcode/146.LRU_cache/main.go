/*
146.LRU 缓存
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；
如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。

注意：函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

思路：实现一个双链表作为LRU结构。
再用一个哈希表方便查找每个节点。
在双向链表的实现中，使用一个伪头部（dummy head）和伪尾部（dummy tail）标记界限，这样在添加节点和删除节点的时候就不需要检查相邻的节点是否存在。
*/
package main

import "fmt"

type DLinkedNode struct {
	Key  int
	Val  int
	Prev *DLinkedNode
	Next *DLinkedNode
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	if capacity <= 0 {
		panic("invalid capacity")
	}
	lc := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	lc.head.Next = lc.tail
	lc.tail.Prev = lc.head
	return lc
}

// 增加到头部
func (lc *LRUCache) addToHead(node *DLinkedNode) {
	node.Next = lc.head.Next
	node.Prev = lc.head
	lc.head.Next.Prev = node
	lc.head.Next = node
}

func (lc *LRUCache) moveToHead(node *DLinkedNode) {
	if lc.head.Next != node {
		lc.removeNode(node)
		lc.addToHead(node)
	}
}

// 移除元素
func (lc *LRUCache) removeNode(node *DLinkedNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// 移除尾部
func (lc *LRUCache) removeTailAndHash() {
	if lc.tail.Prev != lc.head { //由于假头部和假尾部的存在，需要过滤掉没有元素的场景。
		delete(lc.cache, lc.tail.Prev.Key)
		lc.tail.Prev.Prev.Next = lc.tail
		lc.tail.Prev = lc.tail.Prev.Prev

	}
}

func (lc *LRUCache) Get(key int) int {
	if n, ok := lc.cache[key]; !ok {
		return -1
	} else {
		lc.moveToHead(n)
		return n.Val
	}
}

func (lc *LRUCache) Put(key int, value int) {
	if n, ok := lc.cache[key]; !ok { //原本没有，需要增加
		newNode := &DLinkedNode{Key: key, Val: value}
		lc.addToHead(newNode)
		lc.cache[key] = newNode
		if lc.size == lc.capacity { //达到上限的情况，需要移除尾部
			lc.removeTailAndHash()
		} else {
			lc.size++
		}
	} else { //元素存在的情况，更新元素，移动位置
		n.Val = value
		lc.moveToHead(n)
	}
}

func main() {
	lc := Constructor(2)
	lc.Put(1, 0)
	lc.Put(2, 2)
	fmt.Println(lc.Get(1))
	lc.Put(3, 3)
	fmt.Println(lc.Get(2))
	lc.Put(4, 4)
	fmt.Println(lc.Get(1))
	fmt.Println(lc.Get(3))
	fmt.Println(lc.Get(4))
}
