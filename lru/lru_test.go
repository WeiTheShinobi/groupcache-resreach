package lru

import (
	"reflect"
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

func TestCache_Get(t *testing.T) {
	c := New(0, nil)

	c.Add("a", String("1234"))
	v, ok := c.Get("a")
	s := string(v.(String))

	if !reflect.DeepEqual(s, "1234") {
		t.Log("value error")
	}
	t.Log(s)
	t.Log(ok)

	if _, ok = c.Get("c"); ok {
		t.Log("err")
	}
	t.Log(ok)
}

func (c *Cache) Len() int {
	return c.ll.Len()
}

func TestRemoveOldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	c := len(k1 + k2 + v1 + v2)
	lru := New(int64(c), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}
