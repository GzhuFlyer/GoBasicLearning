package main

import (
	"fmt"
	"sync"
)

func main() {
	var i *int
	i = new(int)
	*i = 10
	fmt.Printf("%d\n", *i)

	var u *user
	u = new(user)
	u.name = "你好"
	fmt.Printf("u.name=%s\n", u.name)

	im := make(map[string]int)
	im["hello"] = 3
	im["world"] = 4
	im["hello"] = 5
	fmt.Printf("im=%#v", im)
}

type user struct {
	mut  sync.Mutex
	name string
	age  int
}
