package main

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

func main() {
	lru2, _ := NewLRU(3)
	lru2.Add(1, 2)
	lru2.Add(2, 3)

	lru2.Add(3, 4)

	lru2.Add(4, 5)
	fmt.Println(lru2.Get(4))

	fmt.Println(lru2.Get(1))
	lru2.Get(2)
	lru2.Add(5, 6)
	fmt.Println(lru2.Get(3))
}

type LRU struct {
	lock      sync.RWMutex
	size      int
	evictList *list.List
	items     map[interface{}]*list.Element
}

// entry is used to hold a value in the evictList
type entry struct {
	key   interface{}
	value interface{}
}

// NewLRU constructs an LRU of the given size
func NewLRU(size int) (*LRU, error) {
	if size <= 0 {
		return nil, errors.New("Must provide a positive size")
	}
	c := &LRU{
		size:      size,
		evictList: list.New(),
		items:     make(map[interface{}]*list.Element),
	}
	return c, nil
}

// Add adds a value to the cache.  Returns true if an eviction occurred.
func (c *LRU) Add(key, value interface{}) (evicted bool) {
	// Check for existing item
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		ent.Value.(*entry).value = value
		return false
	}

	// Add new item
	ent := &entry{key, value}
	entry := c.evictList.PushFront(ent)
	c.items[key] = entry

	evict := c.evictList.Len() > c.size
	// Verify size not exceeded
	if evict {
		c.removeOldest()
	}
	return evict
}

// Get looks up a key's value from the cache.
func (c *LRU) Get(key interface{}) (value interface{}, ok bool) {
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		if ent.Value.(*entry) == nil {
			return nil, false
		}
		return ent.Value.(*entry).value, true
	}
	return
}

// Remove removes the provided key from the cache, returning if the
// key was contained.
func (c *LRU) Remove(key interface{}) (present bool) {
	if ent, ok := c.items[key]; ok {
		c.removeElement(ent)
		return true
	}
	return false
}

// RemoveOldest removes the oldest item from the cache.
func (c *LRU) RemoveOldest() (key interface{}, value interface{}, ok bool) {
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
		kv := ent.Value.(*entry)
		return kv.key, kv.value, true
	}
	return nil, nil, false
}

// removeOldest removes the oldest item from the cache.
func (c *LRU) removeOldest() {
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
	}
}

// removeElement is used to remove a given list element from the cache
func (c *LRU) removeElement(e *list.Element) {
	c.evictList.Remove(e)
	kv := e.Value.(*entry)
	delete(c.items, kv.key)
}
