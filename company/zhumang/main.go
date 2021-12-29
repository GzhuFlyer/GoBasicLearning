package main

import (
	"fmt"
	"strings"
	"sync"
)

var connList = make(map[string]int)
var mu1 sync.Mutex
var wg1 = sync.WaitGroup{}

func main() {
	sum := 10
	wg1.Add(sum)
	for i := 0; i < sum; i++ {
		uid := strings.Repeat("a", i+1)
		fd := i
		go func(uid string, fd int) {
			addToConnList(uid, fd)
		}(uid, fd)
	}
	wg1.Wait()
	mu1.Lock()
	fmt.Printf("%+v\n", connList)
	mu1.Unlock() //这里也要获取一下锁
	// 预期输出：map[a:1 aa:2 aaa:3 aaaa:4 aaaaa:5 aaaaaa:6 aaaaaaa:7 aaaaaaaa:8 aaaaaaaaa:9 aaaaaaaaaa:10]
}

func addToConnList(uid string, fd int) {
	mu1.Lock()
	defer mu1.Unlock()
	connList[uid] = fd //判断一下
	wg1.Done()
}
