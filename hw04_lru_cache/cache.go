package hw04lrucache

import (
	"fmt"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	_, ok := lc.items[key]

	if ok {
		lc.items[key].Value = value
		lc.queue.PushFront(lc.items[key])

		return ok
	}

	lc.items[key] = &ListItem{Value: value, CacheKey: key}

	if lc.queue.Len() >= lc.capacity {
		backItem, ok := lc.queue.Back().Value.(*ListItem)
		if ok {
			backItemCacheKey := backItem.CacheKey
			_, ok := lc.items[backItemCacheKey]
			if ok {
				delete(lc.items, backItemCacheKey)
				lc.queue.Remove(lc.queue.Back())
			} else {
				fmt.Println("can't find back item cache key")
			}
		}
	}

	lc.queue.PushFront(lc.items[key])

	return ok
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	listItem, ok := lc.items[key]

	if ok {
		return listItem.Value, ok
	}

	return nil, false
}

func (lc *lruCache) Clear() {
	lc.items = make(map[Key]*ListItem, lc.capacity)
	lc.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
