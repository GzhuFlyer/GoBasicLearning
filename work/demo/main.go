package main

import (
	"errors"
	"fmt"
)

//learn
//https://blog.csdn.net/itcastcpp/article/details/80462619

//str := "fdf"	//error
var str = "dfdf"

const (
//ERR_ELEM = errors.New("fdfd") //error
)

func test1(args ...int) int {
	//fmt.Println(str)
	//a := fmt.Sprintf("abc%d",123)
	//fmt.Println(a)
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	type myInt int
	var i int = 1
	//var j myInt = (myInt)i
	//var j myInt = (myInt)i //error
	var j myInt = myInt(i)
	//var j myInt = i.(myInt)
	fmt.Printf("j=%d\n", j)
	fmt.Println(errors.New("fdfd"))
	return sum
}
func defertest2() {
	var i bool
	//i = 1 //error
	i = true
	//i = bool(1) //error
	i = (1 == 2)
	fmt.Println(i)
	if true {
		defer fmt.Printf("1\n")
	} else {
		defer fmt.Printf("2\n")
	}
	if true {
		defer fmt.Printf("4\n")
	} else {
		defer fmt.Printf("5\n")
	}
	fmt.Printf("3\n")
}

func switchT(in interface{}) {
	switch in {
	case "hello":
		fmt.Println("hello")
		fallthrough
	case 3:
		fmt.Println(3)
	case "world":
		fmt.Println("world")
	}
}

type Slice []byte

func (s *Slice) Remove(value interface{}) error {
	for i, v := range *s {
		if value == v {
			*s = append((*s)[:i], (*s)[i+1])
			return nil
		}
	}
}
func base() {
	//var x = nil //error
	var x interface{} = nil
	var z *string = nil
	//var z string = nil//error
	var y error = nil
<<<<<<< Updated upstream
	fmt.Printf("x = %v, y= %v,z=%v\n",x,y,z)
	//s := make(int)
	//fmt.Println(s)
}

func main()  {
=======
	fmt.Printf("x = %v, y= %v,z=%v\n", x, y, z)
	//s := make(int) //error
	//s := make([]int, 4, 14)
	s := make([]int, 0)
	fmt.Println(cap(s), len(s))
	u := []int{1, 2, 3}
	v := []int{
		1, 2, 3,
		4, 5, 6}
	k := 1
	//q := k++ //error
	//++k  //error
	k++
	fmt.Println(u, v, k)

	//var slice = [4]byte{'a', 'b', 'c', 'd'}
	//Slice.Remove(slice)
}
func main() {
>>>>>>> Stashed changes
	//fmt.Println(test1(1,2,3,4))
	//fmt.Println(test1([]int{1,2,3}...))
	//defertest2()
	//switchT("hello")
	base()
}
