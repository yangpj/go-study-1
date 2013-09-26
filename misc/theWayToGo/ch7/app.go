package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf(" Array at index %d  is %d\n ", i, arr1[i])
	}

	arrTest1()

	arrTest2()

	sliceTest1()

	sliceTest2()

	sliceTest3()

	sliceTest4()

	sliceTest5()

	resliceTest()

	copyAndAppend()
}

func f(a [3]int)   { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }

func arrTest1() {
	var ar [3]int
	f(ar)   // 传递拷贝
	fp(&ar) // 传递引用！
}

func arrTest2() {
	// var arrAge = [5]int{18, 21, 15, 22, 16}
	// var arrLazy = [...]int{5, 6, 7, 8, 22}
	// var arrLazy = []int{5,6,7,8,22}
	var arrKeyValue = [5]string{3: "Cris", 4: "Ron"}
	// var arrKeyValue = []string{3:"Chris",4:"Ron"}

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s \n", i, arrKeyValue[i])
	}
}

func sliceTest1() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5not included!

	// load the array with integers :0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d \n ", i, slice1[i])
	}

	fmt.Printf("The length of the arr1 is %d \n ", len(arr1))
	fmt.Printf("The lenght of slice is %d \n", len(slice1))

}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s

}

func sliceTest2() {

	var arr = [5]int{0, 1, 2, 3, 4}
	fmt.Println(sum(arr[:]))
}

func sliceTest3() {
	var slice1 []int = make([]int, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d \n", i, slice1[i])
	}
}

func sliceTest4() {
	slice1 := make([]int, 4)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for idx, val := range slice1 {
		fmt.Printf("index at %d value is %d \n ", idx, val)
	}

}

func sliceTest5() {
	items := [...]int{10, 20, 30}
	for idx := range items {
		items[idx] = items[idx] * 2
	}

	for idx, val := range items {
		fmt.Printf("item at index %d value is %d \n", idx, val)
	}
}

func resliceTest() {
	slice1 := make([]int, 0, 10)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1] // 不停的改变其长度
		slice1[i] = i
		fmt.Printf("the length of slice is %d \n ", len(slice1))
	}

	// 打印slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d \n ", i, slice1[i])
	}
}

func copyAndAppend() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(n)
	fmt.Printf("copied %d elements\n ", n)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
