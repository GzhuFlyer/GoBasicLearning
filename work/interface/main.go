package main

import "fmt"

type Animal interface {
	Speak() string
}
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "????"
}

type CppProgrammer struct {
}

func (j CppProgrammer) Speak() string {
	return "Design patterns"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, CppProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	var i Animal
	i = Dog{}
	i.Speak()

}

//package main
//
//import "fmt"
//
//type People interface {
//	Speak(string) string
//}
//type Student struct {
//
//}
//
//func (stu *Student)Speak(think string) (talk string) {
//	if think == "bitch"{
//		talk = "You are a good boy"
//	}else {
//		talk = "hi"
//	}
//	return
//}
//
//func main()  {
//	var peo People = Student{}
//	think  := "bitch"
//	fmt.Println(peo.Speak(think))
//}
