package main

import (
	"fmt"
)

func main() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c") //acb
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i) //闭包，地址的引用
		}()
	}
}
