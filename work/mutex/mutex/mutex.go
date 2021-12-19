package main

import (
	"fmt"
	"sync"
)

//output: 没有进行临界区访问控制
//9667
//没有进行临界区访问控制
//9790
var x1 = 0
var mu1 sync.Mutex
var wg1 = sync.WaitGroup{}

func add() {
	mu1.Lock()
	defer mu1.Unlock()
	x1 = x1 + 1
	wg1.Done()
}
func main() {
	sum := 10000
	wg1.Add(sum)
	for i := 0; i < sum; i++ {
		go add()
		//fmt.Println(i)
	}
	wg1.Wait()
	mu1.Lock() //这里也要获取锁，不然有可能变量x1还在缓存中便被打印出来
	fmt.Println("\n---------", x1, "---------")
	mu1.Unlock()
}
