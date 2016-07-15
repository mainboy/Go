package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Println("hello ", name, " I'm ", u.Name)
}

func main() {

	u := User{1, "OK", 12}
	v := reflect.ValueOf(u)

	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("iBoy")}

	mv.Call(args)

	/*
		x := 123
		v := reflect.ValueOf(&x)

		v.Elem().SetInt(956)
		fmt.Println(x)
	*/

}
