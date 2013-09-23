package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	x := Min(1, 3, 2, 0)
	fmt.Printf("The minimum value is : %d\n", x)
	arr := []int{7, 9, 3, 5, 1}
	y := Min(arr...)
	fmt.Printf("the minimum in the array arr is : %d ", y)

	varargs(arr...)
	// defer 测试
	deferTest()
}

func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

// 打印不定长参数示例
func varargs(il ...int) {
	for _, val := range il {
		fmt.Printf("the val is %d \n ", val)
	}
}

func deferTest() {
	trace("deferTest")
	defer untrace("deferTest")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("the val is %d \n ", i)
	}
}

func trace(s string)   { fmt.Println("//....<< entering:", s) }
func untrace(s string) { fmt.Printf("....%s leaving..>>...// \n ", s) }
