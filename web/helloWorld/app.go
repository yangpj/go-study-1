package main

import (
	"fmt"
	http "net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)

	// use the task manager to kill the process !
	//fmt.Println("server is running on 8080")
}
