package initTest

import (
	"fmt"
	"pkg_test/test1"
)

func init() {
	fmt.Println("test -- init")
}

func Test() {
	test1.Test1()
	fmt.Println("this is test.test")
}
