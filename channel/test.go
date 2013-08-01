package main

import "fmt"

func sum(x, y int, c chan int) {
	c <- x + y
}

func main() {
	// init a int type channel
	c := make(chan int)
	go sum(24, 33, c)

	// read the result from channel this will warting for the gorouting
	fmt.Println(<-c)
}
