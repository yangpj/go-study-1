package main

import (
	"fmt"
)

type test struct {
}

type person struct {
	Name string
	Age  int
}

func main() {
	a := test{}

	fmt.Println(a)

	p := person{}
	fmt.Println(p)

	p2 := person{}
	p2.Name = "yiqing"
	p2.Age = 15
	fmt.Println(p2)

	t1()
	// 值拷贝测试！
	t_value_pass(p2)
	fmt.Println(p2)

	// 指针传递
	// 不同于php引用传递 除了函数要声明
	// 调用时仍然需要声明是取地址&
	t_ptr_pass(&p2)
	fmt.Println(p2)

	// 初始化时就取地址！这也是惯用法
	p3 := &person{
		Name: "yiqing2",
		Age:  11,
	}

	t_ptr_pass(p3)
	fmt.Println(p3)

	// 测试匿名结构
	anonymous_struct()

	// 测试嵌套结构
	t3()

	// 测试模拟继承
	t5()
}

func t1() {
	a := person{
		Name: "joe",
		Age:  19,
	}
	fmt.Println(a)
}

/**
* 结构是值拷贝传递 类似php中的数组参数
 */
func t_value_pass(p person) {
	p.Age = 100
	fmt.Println(p)
}

func t_ptr_pass(p *person) {
	p.Age = 13
	fmt.Println(p)
}

/**
* 匿名结构测试 类似java的匿名类
*　　仅用一两次
 */
func anonymous_struct() {
	as := struct {
		k1 string
		k2 string
		k3 int
	}{
		k1: "hello",
		k2: "yes",
		k3: 56,
	}

	fmt.Println(as)
}

// 结构嵌套
type person2 struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

func t3() {
	a := person2{}
	fmt.Println(a)
}

//  测试模拟继承
type Human struct {
	Sex int
}

type Teacher struct {
	Human
	Name string
}
type Student struct {
	Human
	Name string
}

func t5() {
	a := Teacher{
		Name: "yya",
	}
	b := Student{
		Name: "yiqing",
		Human: Human{
			Sex: 1,
		},
	}

	// 内部匿名结构变量修改
	a.Human.Sex = 0
	fmt.Println(a, b)

	// 直接类似yii中behavior的操作
	//  仅当嵌入结构跟外部结构有同名成员变量时
	// 采用 outStruct.innerStructName.member 访问之
	a.Sex = 1
	fmt.Println(a)
}
