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

	//new为引用类型分配内存，并将类型置为该类型的零值。和C的malloc分配内存类似
	var u *user
	u = new(user)
	fmt.Printf("u.name=%s\n", u.name)
	u.name = "你好"
	fmt.Printf("u.name=%s\n", u.name)
	fmt.Printf("u.age=%d\n", u.age)

	//make 一般为map,slice,channel 分配内存
	im := make(map[string]int)
	im["hello"] = 3
	im["world"] = 4
	im["hello"] = 5
	fmt.Printf("im=%#v\n", im)

	//第二个参数指定切片的长度，第三个参数为预留长度，预估切片可能的最大长度
	is := make([]int, 0, 16)
	var tmp_cap float64 = float64(cap(is))
	var tmp_cap_int int = cap(is)
	for i := 0; i < 10300; i++ {
		if len(is) == 1279 {
			fmt.Printf("d")
		}
		is = append(is, 3)
		if 1 < float64(cap(is))/tmp_cap {
			if float64(cap(is))/tmp_cap > 1.25 && float64(cap(is))/tmp_cap < 2 {
				fmt.Printf("%v ;", float64(cap(is))/tmp_cap)
			}
			fmt.Printf("%v ;", float64(cap(is))/tmp_cap)
			fmt.Printf("tmp_cap_int=%v ;", tmp_cap_int)
			fmt.Printf("tmp_cap=%v ;", tmp_cap)
			fmt.Printf("cap(is)=%d\n", cap(is))
		}
		tmp_cap = float64(cap(is))
		tmp_cap_int = cap(is)
	}
	//从打印的参数可以看出，一开始的cap长度是make指定的长度16，后面进行2倍扩容，当容量大于1024的时候以1.25倍的速度进行扩容
	//在 src/runtime/slice.go可以看到相关源码newcap += newcap / 4
}

type user struct {
	mut  sync.Mutex
	name string
	age  int
}
