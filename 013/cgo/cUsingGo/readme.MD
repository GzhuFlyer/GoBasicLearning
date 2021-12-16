#### 通过c静态库，生成mainA可执行文件
````
静态链接方式（-v 是为了看到具体的步骤）：

gcc -v main.c c/cSourceLib/legendtkl.c -I c/cSourceLib/ -o mainA
````

#### 通过c动态库，生成mainLibSoC可执行文件
````    
gcc -v main.c -I ./c/cSourceLib/ -o mainLibSoC ./c/cSourceLib/liblegendtkl.so
````

#### 通过go动态库，生成mainLibSoGo可执行文件
````    
gcc main.c -o mainLibSoGo ./go/legendtkl.so -I ./go/
````