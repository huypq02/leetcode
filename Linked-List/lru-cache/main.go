package main

import (
	"container/list"
	"fmt"
)

type entry struct {
	key, value int
}

type LRUCache struct {
	cap     int
	dll     *list.List
	hashMap map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:     capacity,
		dll:     list.New(),
		hashMap: make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	l, hashMap := this.dll, this.hashMap
	// Check hash map
	if e, ok := hashMap[key]; ok {
		l.MoveToFront(e)
		return e.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.cap < 1 {
		return
	}

	l, hashMap, cap := this.dll, this.hashMap, this.cap

	// Check hash map
	if e, ok := hashMap[key]; ok {
		e.Value.(*entry).value = value
		l.MoveToFront(e)
		return
	}
	if l.Len() >= cap {
		// Remove the back value
		lb := l.Back()
		if lb != nil {
			delete(hashMap, lb.Value.(*entry).key)
			l.Remove(lb)
		}
	}
	pair := &entry{key, value}
	v := l.PushFront(pair)
	hashMap[key] = v
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	cache := Constructor(2)

	cache.Put(1, 1)           // cache is {1=1}
	cache.Put(2, 2)           // cache is {1=1, 2=2}
	fmt.Println(cache.Get(1)) // return 1
	cache.Put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(cache.Get(1)) // return -1 (not found)
	fmt.Println(cache.Get(3)) // return 3
	fmt.Println(cache.Get(4))
}
