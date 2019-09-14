package main

import (
	"fmt"
)

type Power struct {
	age  int
	high int
	name string
}

func main() {

	var i Power = Power{age: 10, high: 178, name: "NewMan"}

	fmt.Printf("type:%T\n", i)
	fmt.Printf("value:%v\n", i)
	fmt.Printf("value+:%+v\n", i)
	fmt.Printf("value#:%#v\n", i)

	fmt.Println("========interface========")
	var interf interface{} = i
	fmt.Printf("%v\n", interf)
	fmt.Println(interf)
}
