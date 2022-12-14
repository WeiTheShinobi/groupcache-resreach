package lru

import (
	"container/list"
)

/*
LRU認為最近被查找過，
之後也會被查找，
淘汰最久沒被查找的
*/

type Cache struct {
	maxBytes  int64
	nBytes    int64
	cache     map[string]*list.Element
	ll        *list.List
	onEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(key string, value Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		cache:     make(map[string]*list.Element),
		ll:        list.New(),
		onEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (Value, bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		en := ele.Value.(*entry)
		return en.value, true
	}
	return nil, false
}

func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)

		en := ele.Value.(*entry)
		delete(c.cache, en.key)
		c.nBytes -= int64(len(en.key)) + int64(en.value.Len())
		if c.onEvicted != nil {
			c.onEvicted(en.key, en.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)

		kv := ele.Value.(*entry)
		kv.value = value
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())

	} else {
		ele := c.ll.PushFront(&entry{
			key:   key,
			value: value,
		})
		c.cache[key] = ele

		c.nBytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.nBytes > c.maxBytes {
		c.RemoveOldest()
	}
}
