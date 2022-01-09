package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const goNum = 10000

var service = os.Args[1]

//go run tcpClient.go localhost:7777
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	wg.Add(goNum)
	timeStart := time.Now()
	for i := 0; i < goNum; i++ {
		go connetServer()
	}
	wg.Wait()
	fmt.Println("spend time: ", time.Now().Sub(timeStart))
	os.Exit(0)
}

func connetServer() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	//_, err = conn.Write([]byte("HEAD / HTTP/1.0 \r\n\r\n"))
	//checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	conn.Close()
	wg.Done()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func ParseIPTest() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
