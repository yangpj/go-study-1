package main

import (
	"fmt"
)

func main() {
	mapTest1()
}

func mapTest1() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 }, // 注意末尾的逗号 没有会报编译错误！
	}
	fmt.Println(mf)
}
