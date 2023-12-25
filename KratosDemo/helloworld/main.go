// @Author zhangjiaozhu 2023/6/13 14:10:00
package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	//multi()
	RW()

}

// 互斥锁
func multi() {

	type safeMap struct {
		Data map[string]int
		sync.Mutex
	}
	//mp := make(map[string]int, 0)
	mp := safeMap{
		Data:  make(map[string]int, 0),
		Mutex: sync.Mutex{},
	}

	list := []string{"A", "B", "C", "D"}

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mp.Mutex.Lock()
			defer mp.Mutex.Unlock()
			for _, v := range list {
				_, ok := mp.Data[v]
				if !ok {
					mp.Data[v] = 0
				}
				mp.Data[v] += 1
			}
		}(i)
	}
	wg.Wait()
	log.Println(mp)
}

// 读写锁

type Cache struct {
	data map[string]string
	sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data:    make(map[string]string, 0),
		RWMutex: sync.RWMutex{},
	}
}
func (c *Cache) Get(key string) string {
	c.RLock()
	defer c.RUnlock()
	value, ok := c.data[key]
	if !ok {
		return ""
	}
	return value
}
func (c *Cache) Set(key, value string) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}

func RW() {
	c := NewCache()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		c.Set("name", "zhang")
	}()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println(c.Get("name"))
		}()
	}
	wg.Wait()
}
