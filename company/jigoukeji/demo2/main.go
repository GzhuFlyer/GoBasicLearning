package main

import "fmt"

func main() {
	test := []byte("AB C D")
	retTest := deleteSpace(test)
	fmt.Println("delete space num=", retTest)
	fmt.Printf("result str=%s\n", test)
}

func deleteSpace(str []byte) int {
	strLen := len(str)
	retLen := 0
	for i := 0; i < strLen; i++ {
		if ' ' != str[i] {
			str[retLen] = str[i]
			retLen++
		}
	}
	for i:=
	return strLen - retLen
}
