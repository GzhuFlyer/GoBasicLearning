package main

import "fmt"

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "name"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "name"}
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	} else {
		fmt.Println("sn1 != sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string //结构体中包含有map,slice是不能比较的
	}{
		age: 12, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{
		age: 12, m: map[string]string{"a": "1"}}
	fmt.Println(sm1, sm2)
	fmt.Println()
}
