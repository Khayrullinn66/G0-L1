package main

import (
	"fmt"
	"sync"
)

type Map struct {
	data map[int]int
	m    sync.RWMutex
}

func NewMap() Map {
	return Map{data: make(map[int]int)}
}

func (c *Map) Add(key, value int, jr *sync.WaitGroup) {
	c.m.Lock()
	c.data[key] = value
	c.m.Unlock()
	jr.Done()
}

func (c *Map) Get(key int) int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.data[key]
}

func (c *Map) PrintAll() {
	c.m.RLock()
	defer c.m.RUnlock()
	for key, value := range c.data {
		fmt.Println(key, " ", value)

	}
}

var N int = 100

func main() {
	mapa := NewMap()
	var jr sync.WaitGroup
	jr.Add(N)
	for i := 0; i <= N; i++ {
		go mapa.Add(i, i*i, &jr)
	}
	jr.Wait()
	mapa.PrintAll()
}
