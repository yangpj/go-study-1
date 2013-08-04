package main

import (
	"fmt"
	"sort"
)

func main() {
	t1()
	t2()
	t3()

	practice()
}

func t1() {
	// key is int and the value is string
	var m map[int]string
	m = map[int]string{}
	fmt.Println(m)

	m2 := make(map[int]string)
	fmt.Println(m2)

	m2[1] = "yes"

	fmt.Println(m2)
	delete(m2, 1)
	a := m2[1]
	fmt.Println(a)

}

func t2() {
	var m map[int]map[string]string
	fmt.Println(m)

	m = make(map[int]map[string]string)
	v, ok := m[1]["key"]
	if ok {
		fmt.Println("exist!")
		fmt.Println(v)

	}
}

func t3() {
	sm := make([]map[int]string, 5)
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "yes"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)

	// map sort
	m2 := map[int]string{5: "yy", 1: "a", 2: "b", 3: "c"}
	fmt.Println(m2)
	s := make([]int, len(m2))
	i := 0
	for k, _ := range m2 {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)
}

func practice() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	m2 := make(map[string]int)

	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m1)
	fmt.Println(m2)
}
