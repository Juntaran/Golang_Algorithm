/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/12/9 20:47
  */

package LRU

import (
	"container/list"
	"github.com/pkg/errors"
)

// CacheNode 结构
type CacheNode struct {
	Key 	interface{}
	Value 	interface{}
}

// 初始化缓存节点
func NewCacheNode(k, v interface{}) *CacheNode {
	return &CacheNode{k, v}
}

// LRU 结构   双链表+map
type LRUCache struct {
	Capacity 	int
	lruList 	*list.List
	cacheMap 	map[interface{}]*list.Element	// list.Element 就是 listnode
}

// 初始化 LRU 结构
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity,
		list.New(),
		make(map[interface{}]*list.Element),
	}
}

// 返回当前 lru 长度
func (lru *LRUCache) Size() int {
	return lru.lruList.Len()
}

// lru 插入值
func (lru *LRUCache) Set(k, v interface{}) error {
	if lru.lruList == nil {
		return errors.New("LRUCache NOT INIT")
	}

	if e, ok := lru.cacheMap[k]; ok {
		// 插入的 key 已经存在，把它挪到顶端
		// 之后更新对应的 value
		lru.lruList.MoveToFront(e)
		e.Value.(*CacheNode).Value = v
		return nil
	}

	// 如果不存在，插入新元素到链表顶端
	newElement := lru.lruList.PushFront(&CacheNode{k, v})
	lru.cacheMap[k] = newElement

	// 判断是否链表已超过 capacity
	// 如果未超过，return
	// 否则删除末端
	if lru.lruList.Len() > lru.Capacity{
		// 双链表移除最后一个
		lastElement := lru.lruList.Back()
		if lastElement == nil {
			return nil
		}
		lru.lruList.Remove(lastElement)

		// 从 map 中删除
		delete(lru.cacheMap, lastElement.Value.(*CacheNode).Key)
	}
	return nil
}

// lru 查询值
func (lru *LRUCache) Get(k interface{}) (v interface{}, ret bool, err error) {
	if lru.cacheMap == nil {
		return v, false, errors.New("LRUCache NOT INIT")
	}

	// 如果 map 中存在该 key
	// 双链表把它移到顶端
	if e, ok := lru.cacheMap[k]; ok {
		lru.lruList.MoveToFront(e)
		return e.Value.(*CacheNode).Value, true, nil
	}

	// map 中未查找到，返回 false
	return v, false, nil
}

// lru 删除
func (lru *LRUCache) Remove(k interface{}) bool {
	if lru.cacheMap == nil {
		return false
	}

	// map 中存在该 key
	// 分别从双链表和 map 中删除
	if e, ok := lru.cacheMap[k]; ok {
		waitToDelete := e.Value.(*CacheNode)
		delete(lru.cacheMap, waitToDelete.Key)
		lru.lruList.Remove(e)
		return true
	}
	return false
}
