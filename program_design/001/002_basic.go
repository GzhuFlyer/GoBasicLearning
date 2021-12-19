package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

//自定义默认路由
type MyMux struct {
}
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path == "/"{
		sayhelloName(w,r)
		return
	}
	http.NotFound(w,r)
	return
}
func sayhelloName(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm() 	//解析参数，默认是不会解析的
	fmt.Println(r.Form)	//这些信息是输出到服务器端的打印信息
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	fmt.Println("---------------")
	for k,v := range r.Form{
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello world!")	//这个写入到w的是输出到客户端的
}
func test1()  {
	http.HandleFunc("/",sayhelloName)	//设置访问的路由
	err := http.ListenAndServe(":9090",nil)	//设置监听的端口
	if err != nil{
		log.Fatal("ListenAndServer: ", err)
	}
}
func test2()  {
	mux := &MyMux{}
	http.ListenAndServe(":9091",mux)
}

func main()  {
	//test1()
	test2()
}