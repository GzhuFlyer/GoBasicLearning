package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("test=", test1(3.0, 4.0))
	//for i := 1; i < 255; i++ {
	//	var j = "http://163.177.151."
	//	var url = j + strconv.Itoa(i)
	//	fmt.Println("url=", url)
	//	resp, err := test2(url)
	//	if err != nil {
	//		log.SetPrefix("get fail: ")
	//		log.SetFlags(0)
	//		log.Printf("%v", err)
	//		//log.Fatalf("Get fail: %v\n", err)
	//		fmt.Fprintf(os.Stderr, "GET fail: %v\n", err)
	//	}
	//	fmt.Println("resp=", resp)
	//}

	respCh := make(chan *http.Response)
	err := make(chan error)
	for i := 1; i < 3; i++ {
		var url = "http://163.177.151." + strconv.Itoa(i)
		go test3(url, respCh, err)
		fmt.Println("url=", url)
	}
	for i := 1; i < 3; i++ {
		fmt.Printf("resp=%v\n", <-respCh)
		fmt.Printf("err=%v\n", <-err)
	}

}

func test1(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func test2(url string) (resp interface{}, err error) {
	resp, err = http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func test3(url string, resp chan *http.Response, err chan error) {
	Tresp, Terr := http.Get(url)
	fmt.Println("url++++++++++++++++++ ", url)
	if err == nil {
		fmt.Println("url==========> ", url)
	}
	if err != nil {
		err <- Terr
		return
	}
	resp <- Tresp
	return
}
