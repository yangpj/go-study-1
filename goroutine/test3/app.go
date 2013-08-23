package main

import (
	"fmt"
	"time"
)

// 声明管道
var ci chan int

func main() {

	ci = make(chan int)

	go ready("Tea", 2)

	go ready("Coffee", 1)

	fmt.Println("i am warting  but not too long !")

	i := 0
	// 用select 读管道
L:
	for {
		select {
		case <-ci:
			i++
			if i > 1 {
				break L
			}

		}
	}
}

func ready(w string, sec int) {

	time.Sleep(time.Duration(sec) * time.Second)

	fmt.Println(w, "is ready")

	// 写管道
	ci <- 2

}
