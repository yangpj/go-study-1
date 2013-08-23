package main

import (
	"fmt"
	"time"
)

func main() {
	go ready("Tea", 2)

	go ready("Coffee", 1)

	fmt.Println("i am warting ...")

	time.Sleep(5 * time.Second)

}

func ready(w string, sec int) {

	time.Sleep(time.Duration(sec) * time.Second)

	fmt.Println(w, "is ready")

}
