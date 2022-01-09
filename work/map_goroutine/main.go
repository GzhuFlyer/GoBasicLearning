package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func main() {
	syncMapTest()
}

func syncMapTest() {
	b := testing.B{}
	BenchmarkSyncMapReadAndWrite(&b)
}

var MapSize = 1000
var LoopCounter = 1000
var MaxGoRoutineNumber = 1000

func BenchmarkSyncMapReadAndWrite(b *testing.B) {
	var m sync.Map

	keys := []string{}
	for i := 0; i < MapSize; i++ {
		keys = append(keys, strconv.Itoa(i))
		m.Store(keys[i], i)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < LoopCounter; i++ {
			if _, ok := m.Load(keys[i]); ok {
				m.Store(keys[i], i+1)
			}
		}
	}
}

func BenchmarkSyncMapReadAndWriteWithMutilGoroutine(b *testing.B) {
	var m sync.Map
	keys := []string{}
	for i := 0; i < MapSize; i++ {
		keys = append(keys, strconv.Itoa(i))
		m.Store(keys[i], i)
	}

	b.ResetTimer()
	wg := sync.WaitGroup{}
	for i := 0; i < MaxGoRoutineNumber; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < b.N; n++ {
				for i := 0; i < LoopCounter; i++ {
					if _, ok := m.Load(keys[i]); ok {
						m.Store(keys[i], i+1)
					}
				}
			}
		}()
	}
	wg.Wait()
}

func mapIsNosaftyIngoroutine() {
	n := 100000
	c := make(map[string]int)
	go func() {
		for j := 0; j < n; j++ {
			c[fmt.Sprintf("%d", j)] = j
		}
	}()
	go func() {
		for j := 0; j < n; j++ {
			fmt.Println(c[fmt.Sprintf("%d", j)])
		}
	}()
	time.Sleep(time.Second * 3)
}
