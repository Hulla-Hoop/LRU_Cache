package lrucache

import (
	"container/list"
	"sync"
	"time"
)

type Item struct {
	Key   string
	Value string
}

type LRUCache struct {
	MU       *sync.Mutex
	Capacity int
	Items    map[string]*list.Element
	List     *list.List
}

func New(n int) *LRUCache {
	var m sync.Mutex
	return &LRUCache{
		Capacity: n,
		Items:    make(map[string]*list.Element),
		List:     list.New(),
		MU:       &m,
	}
}

func (c *LRUCache) Add(key string, value string) bool {
	c.MU.Lock()
	defer c.MU.Unlock()
	// для проверки потокобезопасности
	time.Sleep(1 * time.Second)
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
	c.MU.Lock()
	defer c.MU.Unlock()
	// для проверки потокобезопасности
	time.Sleep(1 * time.Second)
	if data, ok := c.Items[key]; ok {
		c.List.MoveToFront(data)
		return data.Value.(*Item).Value, true
	}
	return "", false
}

func (c *LRUCache) Remove(key string) bool {
	c.MU.Lock()
	defer c.MU.Unlock()
	data, ok := c.Items[key]
	if ok {
		item := c.List.Remove(data).(*Item)
		delete(c.Items, item.Key)
		return true
	}
	return false
}
