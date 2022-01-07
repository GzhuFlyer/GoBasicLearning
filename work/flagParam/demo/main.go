package main

import (
	"flag"
	"fmt"
	"reflect"
)

func main() {
	var name string
	var age int64

	flag.StringVar(&name, "name", "default name", "get name")
	flag.Int64Var(&age, "age", 18, "get age")
	flag.Parse()

	fmt.Println(reflect.TypeOf(name))
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(age))
	fmt.Println(age)
}
