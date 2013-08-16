package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	var a USB
	a = PhoneConnector{"PhoneConnector"}
	a.Connect()

	Disconnect(a)
	Disconnect2(a)
	typeSwitchTest(a)

	typeCastTest()

	emptyInterfaceTest2()
}

func Disconnect(u USB) {
	if pc, ok := u.(PhoneConnector); ok {
		fmt.Println("Disconnected : by PhoneConnector", pc.name)
	}

}

// ok pattern
func Disconnect2(u interface{}) {
	if pc, ok := u.(PhoneConnector); ok {
		fmt.Println("Disconnected : by PhoneConnector", pc.name)
	}

}

func typeSwitchTest(obj interface{}) {

	switch v := obj.(type) {
	case PhoneConnector:
		fmt.Println("the type is PhoneConnector", v.name)
	default:
		fmt.Println("unknown object type")
	}
}

type Connector interface {
	Connect()
}

type USB2 interface {
	Name() string
	//  嵌入接口
	Connector
}

type EmptyInterface interface {
}

func typeCastTest() {
	pc := PhoneConnector{name: "jj"}
	c := Connector(pc)

	c.Connect()
}

func emptyInterfaceTest2() {
	var a interface{}

	fmt.Println(a == nil)
}


