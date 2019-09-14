package main

import (
	"fmt"
)

func main() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer_clousure i = ", i) }()
		fs[i] = func() {
			fmt.Println("closure i = ", i) //闭包使用的时候外层变量内存地址
		}
	}
	for _, f := range fs {
		f()
	}
}
