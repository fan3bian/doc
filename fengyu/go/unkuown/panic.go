package main

import (
	"fmt"
)

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("a")
}
func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover b")
		}
	}()
	panic("b")

	// fmt.Println("b")
}
func C() {
	fmt.Println("c")
}
