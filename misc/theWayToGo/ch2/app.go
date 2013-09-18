package main

// #include <stdlib.h>

import (
	"fmt"
	"runtime"
	// 调用C 库
	"C"
)

func main() {
	fmt.Print("%s", runtime.Version())
	fmt.Println(C.srandom(C.unit(2)))
}
