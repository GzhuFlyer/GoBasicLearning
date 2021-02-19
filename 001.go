package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
)

func test1() {
	var i int = 3
	fmt.Println(i)

	x, y, j := 1, 2, 3 //简短声明，只能在函数内部使用
	fmt.Println(x, y, j)

	_, b := 34, 35
	fmt.Println(b) //_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。将值35赋予b，并同时丢弃34：

	const pi float32 = 3.1415 //常量定于
	const Pi = 3.1415         //第二种定义类型

	var c complex64 = 5 + 5i
	fmt.Printf("value is: %v", c) //支持复数
	fmt.Println()

	var s1 string = "hello" //go中字符串不允许改变
	c1 := []byte(s1)        //将字符串 s1 转换成 []byte 类型
	c1[0] = 'c'
	s2 := string(c1)
	fmt.Printf("%s\n", s2)

	s3 := s1 + s2 //字符串连接操作，+
	fmt.Println(s3)
}

func test2() {
	//分组声明
	const (
		i      = 100
		pi     = 3.14159265
		prefix = "go_"
	)
	var (
		i2      int
		pi2     float32
		prefix2 string
	)
	fmt.Println(i2, pi2, prefix2)
}

func test3() {
	//iota枚举
	const (
		x = iota
		y = iota
		z = iota
		u
	)
	fmt.Println(x, y, z, u)
}

func test4() {
	//array 数组定义
	var arr [10]int //声明了一个int类型的数组
	arr[0] = 42     //数组下标是从0开始的
	arr[1] = 13     //赋值操作
	fmt.Printf("第一个数组数据为：%d\n", arr[0])
	fmt.Printf("最后一个数组数据为：%d\n", arr[9])

	//数组另外一种声明方式
	a := [4]int{1, 2, 3, 4} //声明了一个长度为3的int数组
	b := [10]int{1,2,3}		//声明了一个长度为10的int数组，其中前三个元素初始化为1,2,3；其他默认为0
	fmt.Println(a)
	fmt.Println(b)

	//多维数组
	doubleArray := [2][4] int {[4]int{1,2,3,4},[4]int{5,6,7}}
	fmt.Println(doubleArray)
}

func test5()  {
	//动态数组
	//slice并不是真正意义上的动态数组，而是一个引用类型。
	//slice总是指向一个底层array，slice的声明也可以像 array一样，只是不需要长度。
	x := []byte{'1','2','3'}
	fmt.Println(x)

	//声明一个数组
	var array = [10]byte{'a','b','c','d','e','f','g','h','i','j'}
	//声明两个slice
	var aslice,bslice []byte
	aslice = array[:3]	//等价于aslice = array[0:3]
	aslice = array[5:] //等价于aslice = array[5：10]
	aslice = array[:]	//等价于aslice = array[0:10]
	aslice = array[3:7]	//aSlice包含元素: d,e,f,g，len=4，cap=7
	bslice = aslice[1:3]	//bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
	bslice = aslice[:3]   // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	// 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
	//从概念上面来说slice像一个结构体，这个结构体包含了三个元素： - 一个指针，
	//指向数组中slice指定的开始位 置 - 长度，即slice的长度 - 最大长度，
	//也就是slice开始位置到数组的最后位置的长度
	bslice = aslice[0:5]
	bslice = aslice[:] // bSlice包含所有aSlice的元素: d,e,f,g
	fmt.Println(aslice)
	fmt.Println(bslice)
	/*对于slice有几个有用的内置函数： 
	len 获取slice的长度 cap 获取slice的最大容量 append 向slice里面追加一个或者多个元素，
	然后返回一个和slice一样类型的slice copy 函数copy从源slice的src中复制元素到目标dst，
	并且返回复制的元素的个数*/
}

func test6()  {
	//map的读取和设置也类似slice一样，通过key来操作
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	//var numbers map[string] int
	//另一种map的声明方式
	numbers1 := make(map[string]int)
	numbers1["one"] = 1		//赋值
	numbers1["two"] = 10	//赋值
	numbers1["three"] = 3
	numbers1["three"] = 34

	fmt.Println("第三个数字是：",numbers1["three"])	//读取数据

	/*
	- map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必 须通过key获取
	- map的长度是不固定的，也就是和slice一样，也是一种引用类型
	- 内置的len函数同样适用于 map，返回map拥有的key的数量
	- map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为 one的字典值改为11
	*/

	//通过delete删除map的元素：
	// 初始化一个字典
	rating := map[string]float32{"C":5, "Go":4.5, "python":4.5,"C++":2}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok:= rating["C#"]
	if ok{
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	}else {
		fmt.Println("We have no rating associated with C# in the map")
	}
	delete(rating,"C")

	//map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变：
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	fmt.Println(m)
	m1["Hello"] = "Salut"
	fmt.Println(m)
}

func test7()  {
	//make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。
	/*
	new(T)分配了零值填充的T类型的内存空间，并且返回其 地址，
	即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。
	new返回指针。
	*/
	/*
	内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，
	并且返回一个有初 始值(非零)的T类型，而不是*T。本质来讲，
	导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被 初始化。
	例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；
	在这些项目被 初始化之前，slice为nil。对于slice、map和channel来说，
	make初始化了内部的数据结构，填充适当的值。
	*/
}
func test8()  {
	x := 12
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}
}
func test9()  {
	fmt.Println("test9")
	i := 0
	Here:
		println(i)
		i++
		if i<10 {
			goto Here
		}
}

func test10()  {
	sum := 0;
	for index:=0; index < 10 ; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)
}
func test11()  {
	/*由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错,
	在这种情况下, 可以使用_来丢弃 不需要的返回值
	*/
	//for k,v := range map{
		//fmt.Println("map's key:",k)
		//fmt.Println("map's val:",v)
	//}
	//for _, v := range map{
	//	fmt.Println("map's val:", v)
	//}
}

func test12()  {
	//Go的switch非常灵活，表达式不必是常量或整数
	//Go里面switch默认相当于每个case最后带有break，
	//匹配成功后不会自动向下执行其他case，而是跳出整个switch,
	//但是可以使用fallthrough强制执行后面的case代 码。
	i := 10
	switch  i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2,3,4:		//多值聚合在case里面
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}

	integer := 6
	switch integer {
	case 4:
		fmt.Println("The interger was <= 4")
		fallthrough
	case 5:
		fmt.Println("The interger was <= 5")
		fallthrough
	case 6:
		fmt.Println("The interger was <= 6")
		fallthrough
	case 7:
		fmt.Println("The interger was <= 7")
		fallthrough
	case 8:
		fmt.Println("The interger was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
func max(a,b int) int{
	/*
		函数格式
		func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
			//这里是处理逻辑代码
			//返回多个值 return value1, value2
		}
	*/
	if a>b {
		return a
	}
	return b
}

func test13()  {
	x := 3
	y := 4
	z := 5
	max_xy := max(x,y)
	max_xz := max(x,z)
	fmt.Printf("max(%d,%d)=%d\n",x,y,max_xy)
	fmt.Printf("max(%d,%d)=%d\n",x,z,max_xz)
	fmt.Printf("max(%d,%d)=%d\n",y,z,max(y,z))
}

//func SumAndProduct(A,B int) (int,int) {
//	return A+B,A*B
//}
func SumAndProduct(A,B int) (add int,Multiplied int) {
	add = A+B
	Multiplied = A * B
	return
}

func test14()  {
	//多个返回值测试
	x := 3
	y := 4
	xPLUSY,xTIMESy := SumAndProduct(x,y)
	fmt.Printf("%d + %d = %d\n",x,y,xPLUSY)
	fmt.Printf("%d * %d = %d\n",x,y,xTIMESy)
}

func argFun(arg ...int)  {
	//Go函数支持变参。接受变参的函数是有着不定数量的参数的。
	//为了做到这点，首先需要定义函数使其接受变参：
	for _,n := range arg{
		fmt.Printf("And the number is :%d\n",n)
	}
}

func test15()  {
	//函数变参测试
	argFun(1,2,3,4)
	argFun(5,6)
}
func add1(a *int) int {
	*a = *a + 1
	return  *a
}
func test16()  {
	//GO里面的传值和传指针类似
	x := 3
	fmt.Println("x= ",x)
	x1 := add1(&x)
	fmt.Println("x+1= ", x1)
	fmt.Println("x= ",x)
	//Go语言中string，slice，map这三种类型的实现机制类似指针，
	//所以可以直接传递，而不用取地址后传递 指针。
}
func test17()  {
	//Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。
	//当函数执行到最后时，这些 defer语句会按照逆序执行，最后该函数返回。
	for i:= 0; i<5; i++ {
		defer fmt.Printf("%d\n",i)
	}
	fmt.Println("test17()")
}

type testInt func(int) bool		//声明了一个函数类型
func isOdd(integer int) bool {
	if integer%2 == 0{
		return false
	}
	return true
}
func isEven(integer int) bool {
	if integer%2 == 0{
		return true
	}
	return false
}
func filter(slice []int, f testInt) []int {
	var result [] int
	for _,value := range  slice{
		if f(value) {
			result = append(result,value)
		}
	}
	return result
}
func test18(){
	//函数作为值，类型
	//在Go中函数也是一种变量，我们可以通过type来定义它，
	//它的类型就是所有拥有相同的参数，相同的返回值的一种类型
	slice := []int {1,2,3,4,5,7}
	fmt.Println("slice= ", slice)
	odd := filter(slice,isOdd)
	fmt.Println("Odd elements of slice are: ",odd)
	even := filter(slice,isEven)
	fmt.Println("Even elements of slice are: ", even)
}

var user = os.Getenv("USER")
func test19()  {
	fmt.Println("test19()")
	/*
	Panic 是一个内建函数，可以中断原有的控制流程，
	进入一个令人恐慌的流程中。当函数F调用panic，函数F的执行被中 断，
	但是F中的延迟函数会正常执行，然后F返回到调用它的地方。在调用的地方，
	F的行为就像调用了panic。这一 过程继续向上，直到发生panic的goroutine中所有调用的函数返回，
	此时程序退出。恐慌可以直接调用panic产 生。也可以由运行时错误产生，例如访问越界的数组。
	*/
	if user == ""{
		//panic("no value for $USER")
	}
	/*
	Recover 是一个内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。
	recover仅在延迟函数中有效。在正常 的执行过程中，调用recover会返回nil，
	并且没有其它任何效果。如果当前的goroutine陷入恐慌，调用 recover可以捕获到panic的输入值，
	并且恢复正常的执行。
	*/
}

//struct 类型
//声明一个新的类型
type person struct {
	name string
	age int
}
//比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
//struct也是传值的
func Older(p1,p2 person) (person, int)  {
	if p1.age > p2.age {
		return p1,p1.age-p2.age
	}
	return  p2, p2.age-p1.age
}
type Human struct { name string
	age int
	weight int
}
type Student struct { Human // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}
func test20()  {
	var tom person
	tom.name,tom.age = "Tom",18	//赋值初始化
	//两个字段都写清楚的初始化
	bob := person{age:25,name:"Bob"}
	//按照struct定义顺序初始化值
	paul := person{"Paul",43}
	tb_Older,tb_diff := Older(tom,bob)
	tp_Older,tp_diff := Older(tom,paul)
	bp_Older,bp_diff := Older(bob,paul)
	fmt.Printf("Of %s and %s,%s is older by %d years\n",
		tom.name,bob.name,tb_Older.name,tb_diff)
	fmt.Printf("Of %s and %s,%s is older by %d years\n",
		tom.name,paul.name,tp_Older.name,tp_diff)
	fmt.Printf("Of %s and %s,%s is older by %d years\n",
		bob.name,paul.name,bp_Older.name,bp_diff)

	//struct的匿名字段
	// 我们初始化一个学生
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality) // 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality) // 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age) // 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)
}

type Rectangle struct {
	width,height float64
}

type Circle struct {
	radius float64
}

//面向对象，method(带有接收者的函数)
func (r Rectangle) area() float64 {
	return r.height * r.width
}
//area()方法 是由 Rectangle/Circle 发出的。
func (c Circle) area() float64  {
	return c.radius * c.radius * math.Pi
}
/* 类似于C中的别名，typedef
type months map[string]int
m := months {
	"January":31,
	"February":28, ...
	"December":31,
}
 */
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)
type Color byte
type Box struct {
	width, height,depth float64
	color Color
}
type BoxList [] Box //boxes 的切片
func (b Box) Volume() float64 {	//Volume()定义了接收者为Box，返回Box的容量
	return b.width * b.height * b.depth
}
func (b *Box) setColor(c Color)  {	//SetColor(c Color)，把Box的颜色改为c
	b.color = c
}
func (b1 BoxList) BiggestsColor() Color {
	v := 0.00			//BiggestsColor()定在在BoxList上面，返回list里面容量最大的颜色
	k := Color(WHITE)
	for _,b := range b1 {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}
func (b1 BoxList) PaintItBlack()  {
	for i,_ := range b1{	//PaintItBlack()把BoxList里面所有Box的颜色全部变成黑色
		b1[i].setColor(BLACK)
	}
}
func (c Color) String() string {	//String()定义在Color上面，返回Color的具体颜色(字符串格式)
	strings := []string {"WHITE","BLACK","BLUE","RED","YELLOW"}
	return strings[c]
}
func test21()  {
	fmt.Println("test21()")
	r1 := Rectangle{12,2}
	r2 := Rectangle{9,4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())

	boxes := BoxList{
		Box{4,4,4,RED},
		Box{10,10,1,YELLOW},
		Box{1,1,20,BLACK},
		Box{10,10,1,BLUE},
		Box{10,30,1,WHITE},
		Box{20,20,20,YELLOW},
	}
	fmt.Printf("We have %d boxes in our set\n",len(boxes))
	fmt.Println("The volume of the first on is",boxes[0].Volume(),"cm3")
	fmt.Println("The color of the last one is",boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is ", boxes.BiggestsColor().String())
	fmt.Println("2,The biggest one is ", boxes.BiggestsColor())
	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is",boxes[1].color.String())
	fmt.Println("Obviously,now,the biggest one is", boxes.BiggestsColor().String())
}

type Human1 struct {
	name string
	age int
	phone string
}
type Student1 struct {
	Human1	//匿名字段
	school string
}
//method 继承
type Employee1 struct {
	Human1	//匿名字段
	company string
}
func (h *Human1) SayHi() {
	//method 继承
	fmt.Printf("Hi, I am %s,you can call me on %s\n",h.name,h.phone)
}
//Employee的method重写Human的method
func (e *Employee1) SayHi()  {
	fmt.Printf("Hi,I am %s, I work at %s. Call me on %s\n",e.name,
		e.company, e.phone)//这里可以分割成两行来写。
}
func test22()  {
	//method 继承,重写
	mark := Student1{Human1{"Mark",25,"222-222-YYYY"},"MIT"}
	sam := Employee1{Human1{"Sam",45,"111-888-XXXX"},"Golan Inc"}
	mark.SayHi()
	sam.SayHi()
}
type Human2 struct {
	name string
	age int
	phone string
}
type Student2 struct {
	Human2	//匿名字段
	school string
	loan float32
}
//method 继承
type Employee2 struct {
	Human2	//匿名字段
	company string
	money float32
}
////Human2 对象实现Sayhi方法
//func (h *Human2) SayHi()  {
//	fmt.Printf("Hi, I am %s you can call me on %s\n",h.name, h.phone)
//}
////Huamn2 对象实现Sing方法
//func (h *Human2) Sing(lyrics string) {
//	fmt.Println("La la,la la la ....",lyrics)
//}
////Human2 对象实现Guzzle方法
//func (h *Human2) Guzzle(beerStein string) {
//	fmt.Println("Guzzle Guzzle Guzzle....",beerStein)
//}
////Employee重载Human的Sayhi方法
//func (e *Employee2) SayHi()  {
//	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n",e.name,
//		e.company,e.phone)
//}
////Student2 实现BorrowMoney方法
//func (s *Student2) BorrowMoney(amount float32)  {
//	s.loan += amount
//}
////Employee实现SpendSalary方法
//func (e *Employee2) SpendSalary(amount float32)  {
//	e.money -= amount
//}
////定义interface
//type Men interface {
//	SayHi()
//	Sing(lyrics string)
//	Guzzle(beerStein string)
//}
//type YoungChap interface {
//	SayHi()
//	Sing(song string)
//	BorrowMoney(amount float32)
//}
//type ElderlyGent interface {
//	SayHi()
//	Sing(song string)
//	SpendSalary(amount float32)
//}
//Human2 实现了SayHi方法
func (h Human2) SayHi(){
	fmt.Printf("Hi,I am %s you can call me on %s\n",h.name,h.phone)
}
func (h Human2) Sing(lyrics string) {
	fmt.Println("La la la la...",lyrics)
}
//Employee重载Human2的SayHi方法
func (e Employee2) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n",e.name,
		e.company,e.phone)
}
//Interface Men被Human，Student和Employee实现
//因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}
func test23()  {
	//interface:是一组method的组合
	fmt.Println("test23()")
	mike := Student2{Human2{"Mike",25,"233-222-XXX"},"MIT",0x00}
	paul := Student2{Human2{"Paul",26,"111-222-XXX"},"Harvard",100}
	sam := Employee2{Human2{"Sam",36,"444-222-XXX"},"Golang Inc.",1000}
	Tom := Employee2{Human2{"Sam",36,"444-222-XXX"},"Things Ltd.",5000}

	var i Men	//定义Men类型的变量i

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men,3)
	//T这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0],x[1],x[2] = paul,sam,mike
	for _,value :=range x{
		value.SayHi()
	}
	/*
		空interface:一个函数把interface{}作为参数，
		那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},
		那么也就可以返回任意类型的值。
	*/
}

type Human3 struct {
	name string
	age int
	phone string
}

func (h Human3) String() string {
	return " " +h.name+" - "+strconv.Itoa(h.age)+" years - ✆ " +h.phone+" "
}
func test24()  {
	//interface 函数参数
	/*
		interface的变量可以持有任意实现该interface类型的对象
	*/
	Bob := Human3{"Bob",39,"000-777-xxx"}
	fmt.Println("This Human is : ", Bob)
}
type Element interface { }
type List [] Element
type Person2 struct {
	name string
	age int
}
func (p Person2) String() string {
	return "(name:" + p.name + "- age:" + strconv.Itoa(p.age) + "years)"
}
func test25()  {
	//查看interface变量的存储类型
	//Comma-ok 断言
	list := make(List,3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person2{"Dennis",70}

	for index, element := range list{
		if value, ok := element.(int);ok{
			fmt.Printf("list[%d] is an int and its value is %d\n",index,value)
		}else if value, ok := element.(string);ok{
			fmt.Printf("list[%d] is a string and its valuee is %s\n",index,value)
		}else if value,ok := element.(Person2);ok{
			fmt.Printf("list[%d] is a Person and its value is %s\n",index,value)
		} else {
			fmt.Println("list[%d]is of a different type",index)
		}
	}
}

func test26()  {
	//反射：反射就是动态的获取变量类型信息和值信息的机制
	var a int
	a = 3
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(),typeOfA.Kind(),reflect.ValueOf(a))//分别获取类型名和种类名

	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	v := p.Elem()
	fmt.Println(p,v)
	v.SetFloat(7.1)
	fmt.Println(v)
}
func test27(s string)  {
	//测试gorutine:goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。
	for i:=0; i<5; i++{
		//让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		runtime.Gosched()
		fmt.Println(s)
	}
}
func sum(a []int, c chan int)  {
	sum := 0
	for _,v := range a{
		sum += v
	}
	c <- sum //将sum发送给c
}
func test28()  {
	//测试channel: 必须使用make 创建channel
	//channel可以理解成一个管道
	//ci := make(chan int)
	//cs := make(chan string)
	//cf := make(chan interface{})
	a := []int{7,2,8,-9,4,0}
	c := make(chan int)
	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)
	//x,y := <-c,<-c //从c中接收
	x := <-c
	y := <-c
	//time.Sleep(1) //协程执行的顺序是无序的
	fmt.Println(x,y,x+y)

	//Buffered Channels
	c1 := make(chan int,2)	//可以指定缓冲大小，在这里。设置为0表示无缓冲,设置小于1会出错
	c1 <- 1
	c1 <- 2
	fmt.Println(<-c1)
	fmt.Println(<-c1)
}
func fibonacci(n int, c chan int)  {
	x,y := 1,1
	for i:= 0; i<n; i++{
		c <- x
		x,y = y,x+y
	}
	close(c)
}
func test29()  {
	//Range和Close测试
	fmt.Println("test29()")
	c := make(chan int, 10)
	go fibonacci(cap(c),c)
	for i := range c{
		fmt.Println(i)
	}
}
func main() {
	fmt.Println("hello world")
	test1();test2();test3();test4();test5()
	test6();test7();test8()
	//test9();
	test10()
	test11();test12();test13();test14();test15()
	test16();test17();test18();test19();test20()
	test21();test22();test23();test24();test25()
	test26()
	go test27("world")	//开一个新的Goroutines执行
	test27("hello")	//当前Goroutines执行
	test28();test29()
}
