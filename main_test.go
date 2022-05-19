package main

import (
	"math/rand"
	"sync"
	"testing"
)

// go test -v -bench=. -benchmem
// go test -v -bench=. -benchmem -benchtime=4s
// go test -v -bench=. -benchmem -benchtime=100x

func BenchmarkRwMapStore(b *testing.B) {
	rwMap := NewRWMutexMap()
	for i := 0; i < b.N; i++ {

		rwMap.Store(i, i*i)
	}
}

func BenchmarkSyncMapStore(b *testing.B) {
	var syncMap sync.Map
	for i := 0; i < b.N; i++ {
		syncMap.Store(i, i*i)
	}
}

//
func BenchmarkRwMapLoad(b *testing.B) {
	rwMap := NewRWMutexMap()
	for i := 0; i < 1000000; i++ {
		rwMap.Store(i, i*i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		rwMap.Load(rand.Intn(1000000))
	}
}

func BenchmarkSyncMapLoad(b *testing.B) {
	var syncMap sync.Map
	for i := 0; i < 1000000; i++ {
		syncMap.Store(i, i*i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		syncMap.Load(rand.Intn(1000000))
	}
}

func BenchmarkRwMapDelete(b *testing.B) {
	rwMap := NewRWMutexMap()
	for i := 0; i < 10000; i++ {
		rwMap.Store(i, i*i)
	}
	for i := 0; i < b.N; i++ {
		//rwMap.Delete(i)
		rwMap.Delete(rand.Intn(1000))
	}
}

func BenchmarkSyncMapDelete(b *testing.B) {
	var syncMap sync.Map
	for i := 0; i < 10000; i++ {
		syncMap.Store(i, i*i)
	}
	for i := 0; i < b.N; i++ {
		//syncMap.Delete(i)
		syncMap.Delete(rand.Intn(10000))
	}
}

const elements = 1000000
const findRandElem = elements + elements/100*20
const FrequencyOfStoreItems = 10
const numberOfGoroutines = 4

func BenchmarkRwMapStoreAndLoad(b *testing.B) {
	rwMap := NewRWMutexMap()
	var store []any
	for i := 0; i < elements; i++ {
		rwMap.Store(i, i)
	}

	work := func(ch chan int) {
		for inc := range ch {
			for i := 0; i < 5; i++ {
				if inc%FrequencyOfStoreItems != 0 {
					val, ok := rwMap.Load(rand.Intn(findRandElem))
					if ok {
						store = append(store, val)
					}
				} else if inc%FrequencyOfStoreItems == 0 {
					rwMap.Store(inc+elements, inc)
				}
			}
		}
	}
	c := make(chan int, 4)
	for i := 0; i < numberOfGoroutines; i++ {
		go work(c)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := i; j < i+4; j++ {
			c <- j
		}
	}
}

func BenchmarkSyncMapStoreAndLoad(b *testing.B) {
	var syncMap sync.Map
	var store []any
	for i := 0; i < elements; i++ {
		syncMap.Store(i, i)
	}

	work := func(ch chan int) {
		for inc := range ch {
			for i := 0; i < 5; i++ {
				if inc%FrequencyOfStoreItems != 0 {
					val, ok := syncMap.Load(rand.Intn(findRandElem))
					if ok {
						store = append(store, val)
					}
				} else if inc%FrequencyOfStoreItems == 0 {
					syncMap.Store(inc+elements, inc)
				}
			}
		}

	}
	c := make(chan int, 4)
	for i := 0; i < numberOfGoroutines; i++ {
		go work(c)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := i; j < i+4; j++ {
			c <- j
		}
	}
}
