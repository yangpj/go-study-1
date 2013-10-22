package main

import (
	"fmt"
	"time"
)

func main() {
	go Go()
	time.Sleep(2 * time.Second)

	c_test()
}

func Go() {
	fmt.Println("gO gO go !")
}

func c_test() {
	cn := make(chan bool)
	go func() {
		fmt.Println("this is ni anonymouse funciton ")
		cn <- true
	}()
	<-cn
}
