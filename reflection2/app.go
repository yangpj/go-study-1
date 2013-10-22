package main

import (
	"fmt"
	"reflect"
)

func t1() {
	a := 2
	r := reflect.ValueOf(&a)
	r.Elem().SetInt(133)
	fmt.Println(a)

}

func main() {
	t1()

	u := User{1, "ok", 22}
	SetUser(&u)

	fmt.Println(u)

	// 测试用反射调用对象的某个方法
	t_reflect_method(&u)
}

type User struct {
	Id   int
	Name string
	Age  int
}

func SetUser(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		v = v.Elem()
	}

	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("Bye!")
	}
}

func (u *User) Hello(name string) {
	fmt.Println("Hello ", name, ", my name is ", u.Name)
}

func t_reflect_method(u interface{}) {
	r := reflect.ValueOf(u)
	mv := r.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
}
