package main

import "container/list"

// LRUCache 表示LRU（最近最少使用）缓存
type LRUCache struct {
	cap   int                   // 缓存可以保存的最大项目数
	cache map[int]*list.Element // 哈希表，用于O(1)时间复杂度的缓存项查找
	list  *list.List            // 双向链表，用于维护缓存项的使用顺序
}

// entry 用于在链表中存储键值对(缓存)
type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*list.Element),
		list:  list.New(),
	}
}

// Get 获取给定键的值，如果它存在于缓存中
func (lru *LRUCache) Get(key int) int {
	// 检查键是否存在于缓存中
	if el, ok := lru.cache[key]; ok {
		// 将访问的元素移到列表前端（最近使用）
		lru.list.MoveToFront(el)
		// 返回与键关联的值
		return el.Value.(*entry).value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	// 检查键是否已存在于缓存中
	if el, ok := lru.cache[key]; ok {
		// 如果键已存在,将元素移到链表前端
		lru.list.MoveToFront(el)
		// 更新值
		el.Value.(*entry).value = value
		return
	}

	// 如果缓存已满,需要删除最久未使用的项
	if lru.list.Len() >= lru.cap {
		// 获取链表最后一个元素(最久未使用的)
		oldest := lru.list.Back()
		if oldest != nil {
			// 从列表中移除最旧的元素
			lru.list.Remove(oldest)
			// 从缓存映射中删除相应的条目
			delete(lru.cache, oldest.Value.(*entry).key)
		}
	}

	// 创建新的缓存项
	newEntry := &entry{key, value}
	// 将新项添加到链表前端
	element := lru.list.PushFront(newEntry)
	// 将新条目添加到缓存映射中
	lru.cache[key] = element
}
