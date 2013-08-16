package main

import (
	"bytes"
	"fmt"
	"sync"
)

func ptr_1() {
	var p *int

	fmt.Println("%v", p)

	var i int
	p = &i

	fmt.Println("%v", p)

	*p = 8
	fmt.Println("%v\n", *p)
	fmt.Println("%v\n", i)
}

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func mem_allocate() {

	p := new(SyncedBuffer)
	var v SyncedBuffer

	fmt.Println(p)
	fmt.Println(v)
}

func new_make() {
	var p *[]int = new([]int)

	// make 仅适用于 map slice  channel 并且返回的不是指针
	var v []int = make([]int, 100)

	// new 分配 make 初始化！
	fmt.Println(p)
	fmt.Println(v)
}

type NameAge struct {
	name string
	age  int
}

func struct_test() {
	a := new(NameAge)
	a.name = "Pete"
	a.age = 42

	fmt.Println("%v\n", a)
}

func main() {

	ptr_1()

	// 内存分配
	mem_allocate()

	struct_test()
}
