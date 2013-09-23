package main

import (
	"fmt"
	"runtime"
	"strings"
)

var prompt = "Enter a digit, e.g. 3 “+ “or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { // Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	testIf()
	testFor()
	testForRange()
}

func testIf() {
	if runtime.GOOS == "windows" {
		fmt.Printf("it 's running under windows")
	} else {
		fmt.Println("it run on *nix ")
	}
}

func testFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(strings.Repeat("G", i))
	}
}

func testForRange() {
	str := "Go is a beautifule language !"
	for pos, char := range str {
		fmt.Printf("Character on posiont %d is : %c \n", pos, char)
	}
}
