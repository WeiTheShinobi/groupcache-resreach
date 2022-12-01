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
	m         map[string]*list.Element
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

func (c *Cache) Get(key string) (Value, bool) {
	if v, ok := c.m[key]; ok {
		c.ll.MoveToFront(v)
		en := v.Value.(*entry)
		return en.value, true
	}
	return nil, false
}

func (c *Cache) RemoveOldest() {

}

func (c *Cache) Add(key string, value Value) {

}
