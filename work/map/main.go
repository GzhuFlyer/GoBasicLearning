package main

import "fmt"

type student struct {
	Name string
	Age int
}

func main()  {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou",Age: 23},
		{Name: "li",Age: 42},
		{Name: "wang",Age: 56},
	}
	//for _,stu := range stus{
	//	m[stu.Name] = &stu
	//} //error
	for i:=0; i<len(stus);i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k,v :=range m{
		fmt.Println(k,v)
	}
}
