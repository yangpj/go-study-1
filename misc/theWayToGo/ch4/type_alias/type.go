package main

import (
	"fmt"
)

type TZ int

func main() {
	var a, b TZ = 3, 4
	c := a + b
	// prints : c has the value : 7
	fmt.Printf(" c  has the value : %d ", c)
}
