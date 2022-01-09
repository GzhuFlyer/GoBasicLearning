package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("received back from client: " + reply)
		msg := "received: " + reply
		fmt.Println("sending to client: " + msg)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
