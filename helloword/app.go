package main

import "fmt"

func main() {
	fmt.Println("hello world")

	if 107.17 == (98.6 + 8.57) {
		fmt.Println("equal !")
	} else {
		fmt.Println("not equal!")
	}

    s := 98.6 + 8.57
    if 107.17 == s {
		fmt.Println("equal !")
	} else {
		fmt.Println("not equal!")
	} 

}
