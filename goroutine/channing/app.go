package main

import (
	"flag"
	"fmt"
)

var gorutionNum = flag.Int("n", 1000, "please input the num of gorutine you want to run ")

func f(lft, rgt chan int) {
	lft <- 1 + <-rgt
}

func main() {
	flag.Parse()
	lftMost := make(chan int)
	var left, right chan int = nil, lftMost

	for i := 0; i < *gorutionNum; i++ {
		left, right = right, make(chan int)

		go f(left, right)
	}
	right <- 0     // start the chaning
	x := <-lftMost // wait for completion
	fmt.Println(x) // 10000 ,approx . 1.5   s

}
