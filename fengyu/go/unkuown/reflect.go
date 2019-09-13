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
type Manager struct {
	User
	title string
}

func (u User) Hello() {
	fmt.Println("Hello world.")
}
func main() {
	u := User{1, "OK", 12}
	// Info(u)
	m := Manager{User: User{1, "OK", 12}, title: "123"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))

	a := 123
	v := reflect.ValueOf(&a)
	v.Elem().SetInt(999)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", v)

	Set(&u)
	fmt.Println(u)

	u.Echo("zhangshuyi")

	value := reflect.ValueOf(u)
	method := value.MethodByName("Echo")
	args := []reflect.Value{reflect.ValueOf("quao")}
	method.Call(args)
}
func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}
	// if f := v.FieldByName("Name"); f.Kind() == reflect.String {
	// 	f.SetString("BYEBYE")
	// }
	f := v.FieldByName("Name1")
	if !f.IsValid() {
		fmt.Println("BAD")
	}
	if f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	count := t.NumField()
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("XX")
		return
	}
	for i := 0; i < count; i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%5s:%v = %v\n", f.Name, f.Type, val) //%num : 如果不足num，则用空格补全
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%5s:%v\n", m.Name, m.Type) //%num : 如果不足num，则用空格补全

	}
}

func (u User) Echo(str string) {
	fmt.Println("you say:", str, "my name is", u.Name)
}
