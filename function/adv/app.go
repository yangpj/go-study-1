package main 

import (
	"fmt"
)

Brawndo := func(x string) func(string) {
	return func(y string){
		fmt.Println("BRAWNDO , THE ",x,y)
	}
}


func main() {
	Brawndo("THIRST")("MULTLATOR")
}


