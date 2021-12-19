package main

<<<<<<< Updated upstream
import (
	"fmt"
	"runtime"
)

type panicContext struct {
	function string
}

func ProtectRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时出错
			fmt.Println("运行时出错:", err)
		default: //非运行时出错
			fmt.Println("非运行时出错.:", err)
		}
	}()
	entry()
}
func main() {
	fmt.Println("运行前")
	ProtectRun(func() {
		fmt.Println("手动宕机前")
		panic(&panicContext{
			"手动触发panic",
		})
		fmt.Println("手动宕机后")
	})
	ProtectRun(func() {
		fmt.Println("复制宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})
	fmt.Println("运行后")
=======
import "fmt"

func main() {
	fmt.Printf("hello")
>>>>>>> Stashed changes
}
