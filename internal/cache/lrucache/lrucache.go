package lrucache

import (
	"container/list"
)

type Item struct {
	Key   string
	Value string
}

type LRUCache struct {
	Capacity int
	Items    map[string]*list.Element
	List     *list.List
}

func New(n int) *LRUCache {
	return &LRUCache{
		Capacity: n,
		Items:    make(map[string]*list.Element),
		List:     list.New(),
	}
}

func (c *LRUCache) Add(key string, value string) bool {
	if data, ok := c.Items[key]; ok {
		c.List.MoveToFront(data)
		data.Value.(*Item).Value = value
		return false
	}
	if c.List.Len() == c.Capacity {
		if element := c.List.Back(); element != nil {
			item := c.List.Remove(element).(*Item)
			delete(c.Items, item.Key)
		}
	}
	item := &Item{
		Key:   key,
		Value: value,
	}
	data := c.List.PushFront(item)
	c.Items[item.Key] = data
	return true
}

func (c *LRUCache) Get(key string) (string, bool) {
	if data, ok := c.Items[key]; ok {
		c.List.MoveToFront(data)
		return data.Value.(*Item).Value, true
	}
	return "", false
}

func (c *LRUCache) Remove(key string) bool {
	data, ok := c.Items[key]
	if ok {
		item := c.List.Remove(data).(*Item)
		delete(c.Items, item.Key)
		return true
	}
	return false
}
