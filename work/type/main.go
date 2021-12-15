package main

import "fmt"

type myInt1 int
type myInt2 = int

type User struct {
}
type MyUser1 User
type MyUser2 = User

func (i MyUser1) m1() {
	fmt.Println("MyUser1.m1")
}
func (i User) m2() {
	fmt.Println("User.m2")
}

func main() {
	var u1 MyUser1 //MyUser1 是定义的新的类型了
	var u2 MyUser2
	u1.m1()
	u2.m2()

	var i int = 0
	var i1 myInt1 = myInt1(i) //创建的新类型不能直接赋值，需要转换
	var i2 myInt2 = i         //别名，可以直接赋值
	fmt.Println(i1, i2)
	typeInterface()
}

func typeInterface() {
	i := GetValue()
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case rune:
		fmt.Println("rune")
	case interface{}:
		fmt.Println("interface")
	default:
		fmt.Println("unknown")
	}
}

func GetValue() interface{} {
	return 1
}

//func GetValue() int {	//error ,type只在interface里面有
//	return 'd'
//}
