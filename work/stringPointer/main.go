package main

import (
	"fmt"
)

func main() {
	s1 := "王总好。《西虹市首富》"
	s2 := "good morning"
	s1p := &s1
	s2p := &s2
	fmt.Println("s1=", *s1p)
	fmt.Println("s2=", *s2p)

	s1Byte := []byte(s1)
	for i := 0; i < len(s1Byte); i++ {
		fmt.Printf("s1(%d)=%c ", i, s1Byte[i])
	}
	fmt.Println("\n=======================")
	s1Rune := []rune(s1)
	for i := 0; i < len(s1Rune); i++ {
		fmt.Printf("s1(%d)=%c ", i, s1Rune[i])
	}
	fmt.Println("\n=======================")
	s2Byte := []byte(s2)
	for i := 0; i < len(s2Byte); i++ {
		fmt.Printf("s2(%d)=%c ", i, s2Byte[i])
	}
	fmt.Println("\n=======================")
	s2Rune := []rune(s2)
	for i := 0; i < len(s2Rune); i++ {
		fmt.Printf("s2(%d)=%c ", i, s2Rune[i])
	}
}
