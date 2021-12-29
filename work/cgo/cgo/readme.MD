##参考：
####https://blog.csdn.net/fuyuande/article/details/89178640
####https://www.sohu.com/a/451319375_100093134
####https://github.com/golang/go/wiki/cgo
知识点：
1，通过 import “C” 语句启用 CGO 特性后，CGO 会将上一行代码所处注释块的内容视为 C 代码块，被称为 序文（preamble）。
2，gcc -c hello.c  
3，ar -cru libhello.a hello.o
