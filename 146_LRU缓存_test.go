package Code

import (
	"container/list"
	"sync"
	"testing"
)

/**
请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value)如果关键字key 已经存在，则变更其数据值value ；如果不存在，则向缓存中插入该组key-value 。如果插入操作导致关键字数量超过capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lru-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test146(t *testing.T) {
	lru := Constructor146(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	t.Log(lru.Get(1)) //1
	lru.Put(3, 3)
	t.Log(lru.Get(2)) //-1
	lru.Put(4, 4)
	t.Log(lru.Get(1)) //-1
	t.Log(lru.Get(3)) //3
	t.Log(lru.Get(4)) //4
}

type LRUCache struct {
	cache     map[int]*list.Element
	cacheList *list.List
	lock      *sync.Mutex
	total     int
	cap       int
}

type entry struct {
	key   int
	value int
}

func Constructor146(capacity int) LRUCache {
	return LRUCache{
		cache:     make(map[int]*list.Element, capacity),
		cacheList: list.New(),
		lock:      &sync.Mutex{},
		total:     0,
		cap:       capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	this.lock.Lock()
	defer this.lock.Unlock()

	value, ok := this.cache[key]
	if ok {
		this.cacheList.MoveToFront(value)
		if value.Value.(*entry) == nil {
			return -1
		}
		return value.Value.(*entry).value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	this.lock.Lock()
	defer this.lock.Unlock()

	vPtr, ok := this.cache[key]
	if ok {
		this.cacheList.MoveToFront(vPtr)
		vPtr.Value.(*entry).value = value
		return
	}

	if this.total == this.cap {
		ele := this.cacheList.Back()
		delete(this.cache, ele.Value.(*entry).key)
		this.cacheList.Remove(ele)
		this.total--
	}
	newElem := &entry{key: key, value: value}
	newEntry := this.cacheList.PushFront(newElem)
	this.cache[key] = newEntry
	this.total++
}

type LRUCacheV2 struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func ConstructorV2(capacity int) LRUCacheV2 {
	l := LRUCacheV2{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCacheV2) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCacheV2) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCacheV2) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCacheV2) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCacheV2) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCacheV2) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
