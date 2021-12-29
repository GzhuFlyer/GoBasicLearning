package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w, r)
		return
	}
	if r.URL.Path == "/good" {
		goodMorning(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprint(w, "hello world")
}
func goodMorning(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "good morning")
}

func main() {
	//ListenAndServe函数传递的第二个参数不为nil，则使用第二个参数设置的路由规则，源码可以在http包里面看到，为nil默认调用DefaultServeMux
	//func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
	//	handler := sh.srv.Handler
	//	if handler == nil {
	//		handler = DefaultServeMux
	//	}
	//	if req.RequestURI == "*" && req.Method == "OPTIONS" {
	//		handler = globalOptionsHandler{}
	//	}
	//	handler.ServeHTTP(rw, req)
	//}
	//http.HandleFunc("/", sayHello)
	//http.HandleFunc("/good", goodMorning)
	mux := &MyMux{}
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
