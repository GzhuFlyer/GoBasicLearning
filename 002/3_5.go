package main

import (
	"crypto/sha256"
	. "fmt"
	"unicode/utf8"
)

func main() {
	//string_test()
	//array_test()
	//sha_test()
	//slice_test()
	map_test()
}

func string_test()  {
	s := "Hello, 世界"
	Println(len(s))
	Println(utf8.RuneCountInString(s))
	Printf("s=%s\n", s)
	Printf(string(65))
	Printf(string(0x4eac))		//将十六进制转换成字符
	Printf(string(0x123456))	//位置字符会打印出问号
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])	//utf8 解码
		Printf("%d\t%c\t", i, r)
		i += size
	}
	Println("")
	for i, r := range s {
		Printf("%d\t%q\t%d\n", i, r, r)
	}
}
func array_test()  {
	var a [3]int = [...]int{1,2,3}
	var b [3] int
	var q [3]int = [3]int{1,2,3}
	var r [3]int = [3]int{1,2}
	Println(r[1])
	Println(q)
	for i,v := range a{
		Printf("a[%d]=%v\t",i,v)
	}
	clear_array_test(&a)
	for i,v := range a{
		Printf("a[%d]=%v\t",i,v)
	}
	Println(a==b,a==q)
	for i,v := range b{
		Printf("b[%d]=%v\t",i,v)	//默认值为0
	}
	Println("")
	type Current int
	const (
		USD Current  = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD:"美元",EUR:"欧元",GBP:"英镑",RMB:"人民币"}
	Println(RMB,symbol[RMB])
	north := symbol[1:]
	Println(north)
}

func clear_array_test(ptr *[3]int)  {
	*ptr = [3]int{}
}

func sha_test()  {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	Printf("%x\t %x\t %t \t %T\n",c1,c2,c1 == c2,c1)
}

func slice_test()  {
	var runes []rune
	for _,r := range "hello, 世界"{
		runes = append(runes,r)
	}
	Printf("%q\n",runes)

	var x,y []int
	for i:=0; i<10; i++{
		y = appendInt(x,i)
		//y = append(x,i)
		Printf("%d cap=%d\t%v\n",i,cap(y),y)
		x = y
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x){
		z = x[:zlen]
	}else {
		zcap := zlen
		if zcap < 2*len(x){
			zcap = 2*len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z,x)
	}
	z[len(x)] = y
	return  z
}

func map_test()  {
	//ages := make(map[string]int)
	//ages["alice"] = 31
	//ages["charlie"] = 34
	ages := map[string]int{
		"alice": 31,
		"charlie": 34,
	}
	Println(ages)
	Println(ages["alice"])
	delete(ages,"alice")
	Println(ages)
	Println(ages["alice"])
	delete(ages,"alice")
	ages["bob"]++
	ages["dad"]++
	Println(ages)
	Println(ages["alice"])
	//查找ages里面是否有键为“bob”的值，有的话ok为1
	if age,ok := ages["bob"]; ok{Printf("bob %v\n",age)}
	if age,ok := ages["frank"]; ok{Printf("frank %v\n",age)}
	println(equal(map[string]int{"A":0}, map[string]int{"B":42}))
	println(equal(map[string]int{"A":0}, map[string]int{"A":42}))
	println(equal(map[string]int{"A":0}, map[string]int{"A":0}))
}

func equal(x,y map[string]int) bool {
	if len(x) != len(y){
		return false
	}
	//遍历 x 中的值，放在k中，如果 y 中没有键为 k 的值，或者y中键k的值y不为xv,则返回false
	for k,xv := range x{
		if yv,ok := y[k]; !ok || yv != xv{
			return false
		}
	}
	return true
}