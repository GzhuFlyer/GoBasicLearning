package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":2222"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	udpCheckErr(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	//for {
	go handleUdpClient(conn)
	//}
	for {

	}
}

func handleUdpClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}

func udpCheckErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
