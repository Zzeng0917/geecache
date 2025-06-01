package singleflight

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
} // call 代表正在进行中

type Group struct {
	mu sync.Mutex
	m  map[string]*call
} // Group 主数据结构

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()         //如果请求正在进行，则等待
		return c.val, c.err // 请求结束，返回结果
	}
	c := new(call)
	c.wg.Add(1)  //发起请求前加锁
	g.m[key] = c //添加到g.m中，表明key已经有对应的请求在处理
	g.mu.Unlock()

	c.val, c.err = fn() //调用fn，发起请求
	c.wg.Done()         //请求结束，解锁

	g.mu.Lock()
	delete(g.m, key) //更新g.m，删除key对应的请求
	g.mu.Unlock()

	return c.val, c.err
} // 针对相同的key，只会执行一次fn，并返回结果
