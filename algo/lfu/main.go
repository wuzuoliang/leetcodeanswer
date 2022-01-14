package main

import "fmt"

type Node struct {
	Key   int
	Value int
	freq  int
	pre   *Node
	next  *Node
}

type LFUCache struct {
	cap     int
	HashMap map[int]*Node
	head    *Node
	tail    *Node
}

func LFUConstructor(capacity int) LFUCache {
	lfuCache := LFUCache{}
	lfuCache.cap = capacity
	lfuCache.HashMap = make(map[int]*Node, capacity)
	return lfuCache
}

func (l *LFUCache) Get(key int) int {
	if value, ok := l.HashMap[key]; ok {
		value.freq++
		l.refreshNode(value)
		return value.Value
	} else {
		return -1
	}
}

func (l *LFUCache) Put(key, value int) {
	if v, ok := l.HashMap[key]; !ok {
		if len(l.HashMap) >= l.cap {
			oldKey := l.removeNode(l.head)
			delete(l.HashMap, oldKey)
		}
		node := Node{Key: key, Value: value, freq: 1}
		l.addNode(&node)
		l.HashMap[key] = &node
	} else {
		v.Value = value
		v.freq++
		l.refreshNode(v)
	}
}

func (l *LFUCache) refreshNode(node *Node) {
	if node == l.tail {
		return
	}
	l.removeNode(node)
	l.addNode(node)
}

func (l *LFUCache) removeNode(node *Node) int {
	if node == l.tail {
		l.tail = l.tail.pre
		l.tail.next = nil
	} else if node == l.head {
		l.head = l.head.next
		l.head.pre = nil
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	return node.Key
}

func (l *LFUCache) addNode(node *Node) {
	if l.head == nil && l.tail == nil {
		l.head = node
		l.tail = node
		return
	}
	head := l.head
	for head != nil && node.freq >= head.freq {
		head = head.next
	}
	if head == nil {
		l.tail.next = node
		node.pre = l.tail
		l.tail = node
	}
	if head != nil {
		head.pre.next = node
		node.pre = head.pre
		head.pre = node
		node.next = head
	}

	l.head.pre = nil
	l.tail.next = nil
	return
}

func main() {
	lfuCache := LFUConstructor(3)
	lfuCache.Put(1, 3)
	lfuCache.Put(2, 4)
	lfuCache.Put(3, 5)
	fmt.Println(lfuCache) // 此时的链表顺序应该是 1 : 3   2 : 4   3 : 5   l.head 为1 : 3   l.tail 为3 : 5
	lfuCache.Get(1)
	fmt.Println(lfuCache) // 此时的链表顺序应该是  2 : 4   3 : 5  1 : 3   l.head 为2 : 4    l.tail 为1 : 3
	lfuCache.Put(4, 6)
	fmt.Println(lfuCache) // 此时的链表顺序应该是  3 : 5  4 : 6   1 : 3    l.head 为3 : 5   l.tail 为1 : 3

}
