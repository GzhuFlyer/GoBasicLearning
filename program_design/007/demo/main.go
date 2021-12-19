package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	test1()
	test2()
}

func test1() {
	var w io.Writer
	w = os.Stdout //*os.File 有Write方法
	w = new(bytes.Buffer)
	//w = time.Second
	println(w)
}

func test2() {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	fmt.Println(any)
	type Video struct {
		//Stream()	(io.ReadCloser,error)
		//Format()	string
	}
}
