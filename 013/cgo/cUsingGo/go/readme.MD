###链接：https://zhuanlan.zhihu.com/p/32066522
````
构建 so 文件。Go 构建出来的 so 文件比较大，因为 Go 有 runtime。

$ go build -o legendtkl.so -buildmode=c-shared legendtkl.go
````