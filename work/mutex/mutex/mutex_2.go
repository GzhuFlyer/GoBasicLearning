package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0
var mu sync.Mutex
var wg = sync.WaitGroup{}

func add1() {
	mu.Lock()
	x = x + 1
	mu.Unlock()
	wg.Done()
}
func main() {
	sum := 10000
	wg.Add(sum)
	for i := 0; i < sum; i++ {
		go add1()
	}
	time.Sleep(time.Second) //这里加延时，不然最后打印结果不是 sum的值10000，不是预期值
	fmt.Println("\n---------", x, "---------")
	wg.Wait()
}
