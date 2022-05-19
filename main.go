package main

import (
	"sync"
)

func main() {

}

type RWMutexMap struct {
	mu sync.RWMutex
	m  map[int]int
}

func NewRWMutexMap() *RWMutexMap {
	return &RWMutexMap{
		m: make(map[int]int, 100000),
	}
}

func (c *RWMutexMap) Load(key int) (int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *RWMutexMap) Store(key int, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

func (c *RWMutexMap) Delete(key int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
}
