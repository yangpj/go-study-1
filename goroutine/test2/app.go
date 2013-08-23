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

	//  两次读
	<-ci
	<-ci
}

func ready(w string, sec int) {

	time.Sleep(time.Duration(sec) * time.Second)

	fmt.Println(w, "is ready")

	// 写管道
	ci <- 2

}
