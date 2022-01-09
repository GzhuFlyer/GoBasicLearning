package main

import (
	"fmt"
	"os"
)

func operatorDirectory() {
	os.Mkdir("good", 0777)
	os.MkdirAll("good/morning/frank", 0777)
	err := os.Remove("good")
	if err != nil {
		fmt.Println(err)
	}
	err = os.RemoveAll("good")
}
func writeFile() {
	userFile := "testFile.txt"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	for i := 0; i < 10; i++ {
		fout.WriteString("hello world!\n")
		fout.Write([]byte("good morning!\r\n"))
	}
}
func readFile() {
	userFile := "testFile.txt"
	f1, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
	}
	buf := make([]byte, 1024)
	for {
		n, _ := f1.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
func main() {
	operatorDirectory()
	writeFile()
	readFile()
}
