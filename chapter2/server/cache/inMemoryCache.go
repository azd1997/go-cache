package cache

import "sync"

// inMemoryCache 内存缓存
// 包内可见，外部不可见

type inMemoryCache struct {
	c     map[string][]byte // 保存键值对
	mutex sync.RWMutex      // 读写锁，对map的并发访问提供读写锁保护  mutex 互斥
	Stat                    // 记录缓存状态		内嵌组合。 类似于 继承自Stat类
}
// 实现cache接口Set方法
func (c *inMemoryCache) Set(k string, v []byte) error {
	c.mutex.Lock()	//上写锁   写操作或者又读又写操作必须上锁，但读操作可以多个goroutine同时读map
	defer c.mutex.Unlock()

	tmp, exist := c.c[k]
	if exist {
		c.del(k, tmp)
	}

	c.c[k] = v
	c.add(k, v)

	return nil
}

// 实现Cache接口Get方法
func (c *inMemoryCache) Get(k string) ([]byte, error) {
	c.mutex.RLock()	//读锁
	defer c.mutex.RUnlock()

	return c.c[k], nil
}

// 实现Cache接口Del方法
func (c *inMemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	v, exist := c.c[k]
	if exist {
		delete(c.c, k)
		c.del(k, v)
	}
	return nil
}

// 实现Cache接口GetStat方法
func (c *inMemoryCache) GetStat() Stat {
	return c.Stat
}

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{make(map[string][]byte), sync.RWMutex{}, Stat{}}
}