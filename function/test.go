package main

import (
	"fmt"
)

func main() {
	max_test(2, 3, 444, 445)

	s1 := []int{1, 2, 3}
	t_slice(s1)

	i1 := 3
	p_test(&i1)

	// 函数名作为变量赋值
	f1 := func_as_var
	f1()

	af := func() {
		fmt.Println("this is anonymouse function mindjet")
	}
	af()

	// 匿名带直接调用
	func() {
		fmt.Println("this is anomymouse function declration and invoke after defined")
	}()

	// 闭包测试
	cf := closure(10)
	fmt.Println(cf(4))
	fmt.Println(cf(2))
	// defer 测试 可实现finally ，析构函数的特征
	deferTest()

	// 错误恢复机制的测试
	exception_test()
}

func A(a int, b string) int {
	return 2
}

func B() (a, b, c int) {
	a, b, c = 1, 3, 5
	return a, b, c
}

/**
* 实际变为slice的拷贝
 */
func max_test(a ...int) {
	fmt.Println(a)
}

/**
* 内部会改变slice 的某值
 */
func t_slice(s []int) {
	s[0] = 100

	fmt.Println(s)
}

/**
* 指针做参
 */
func p_test(a *int) {
	// 用指针来修改变量对应的值
	*a = 399
	fmt.Println(*a)
}

func func_as_var() {
	fmt.Println("in func_as_var")
}

func closure(a int) func(int) int {

	return func(b int) int {

		return a + b
	}
}

func deferTest() {
	a, b, c := "a", "b", "c"
	defer fmt.Println(a)
	defer fmt.Println(b)
	defer fmt.Println(c)

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)

		//  闭包
		defer func() {
			fmt.Println(i)
		}()
	}
}

/**
* go 中没有异常
* 用panic 跟recover来做
 */
func exception_test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("here do recover operation!")
		}
	}()

	panic("this is like a exception")

}
