package lru

import "container/list"

type Cache struct {
	maxBytes  int64      //允许使用的最大内存
	nbytes    int64      //当前已使用的内存
	ll        *list.List //双向链表
	cache     map[string]*list.Element
	onEvicted func(key string, value Value) //某条记录被移除时的回调函数，可以为nil
}

// 双向链表的数据类型
type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		onEvicted: onEvicted,
	}
}

// 在字典中找到对应的双向链表的节点，然后将其移到队尾
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// remove
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back() //取出队首元素
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len()) //更新已使用内存
		if c.onEvicted != nil {
			c.onEvicted(kv.key, kv.value)
		}
	}
}

// add and write
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	//如果超过了最大的设定值，则移除最少的访问节点
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// 返回值所占用的内存大小
func (c *Cache) Len() int {
	return c.ll.Len()
}
