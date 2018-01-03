package link

import "container/list"

type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

type Cache struct {
	MaxEntries int // max entries
	OnEvicted  func(key Key, value interface{})
	list       *list.List
	cache      map[interface{}]*list.Element
}

// create a Cache
func New(maxEntries int) *Cache {
	return &Cache{
		MaxEntries: maxEntries,
		list:       list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

// insert a map to cache
func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.list = list.New()
	}
	if ee, ok := c.cache[key]; ok {
		c.list.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}
	ele := c.list.PushFront(&entry{key, value})
	c.cache[key] = ele
	if c.MaxEntries != 0 && c.list.Len() > c.MaxEntries {
		c.RemoveOldest()
	}
}

// get the cache by key
func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.list.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return
}

// remove an element by key
func (c *Cache) Remove(key Key) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
	}
}

// delete oldest
func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	ele := c.list.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

// get thc cache size
func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.list.Len()
}

// remove the element
func (c *Cache) removeElement(e *list.Element) {
	c.list.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cache, kv.key)
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}
