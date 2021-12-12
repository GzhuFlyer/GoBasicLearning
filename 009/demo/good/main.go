package main

import (
	hello "demo/internal/world"
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("NumCPU: %d, GOMAXPROCS: %d\n", runtime.NumCPU(), runtime.GOMAXPROCS(1))
	hello.Add(3, 4)
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
