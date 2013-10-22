package main

import (
	"fmt"
	"reflect"
	"strings"
)

type TagType struct {
	field1 bool   "An important answer"
	field2 string "The name of the thing "
	field3 int    "How much there are"
}

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type Person struct {
	firstName string
	lastName  string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf("The int is %d \n ", ms.i1)
	fmt.Printf("The float is %f \n", ms.f1)
	fmt.Printf("The string is %s \n", ms.str)
	fmt.Print(ms)
	// 用方法来处理结构体

	testOpOnStruct()

	// 反射tag
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}

}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func testOpOnStruct() {
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s \n", pers1.firstName, pers1.lastName)

	// 2-struct as a pointer
	pers2 := new(Person)
	pers2.firstName = "chris"
	pers2.lastName = "Woodward"
	//(*pers1).lastName = "Woodward" // 仍旧是有效语法
	fmt.Printf("The name of the person is %s %s \n ", pers2.firstName, pers2.lastName)

	// 3-struct as literal
	pers3 := &Person{"Chris ", "Woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s \n", pers3.firstName, pers3.lastName)
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
