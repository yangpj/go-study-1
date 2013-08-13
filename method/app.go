package main

import (
	"fmt"
)

// 声明底层类型的别名
type MyInt int8

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
}

func main() {
	a := A{}
	a.Print()

	b := B{Name: "b"}
	b.Print()

	// 基类型别名的自定义方法
	var ma MyInt
	ma = 1
	// method value
	ma.Print()
	// method expression 另一种调用形式
	(MyInt).Print(ma)

	// 递增练习
	var ma2 MyInt
	ma2.increase()
	ma2.Print()
	// 第二次的奇怪现象
	ma2.increase()
	ma2.Print()
}

func (a A) Print() {
	fmt.Println(a)
}

// 注意传值 跟引用类型传递的区别
func (b *B) Print() {
	fmt.Println(b)
}

func (ma MyInt) Print() {

	fmt.Println(ma)
}

func (ma *MyInt) increase() {
	// 类型转换！
	*ma += MyInt(100)
}
