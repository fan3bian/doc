package main

import (
	"fmt"
)

func main() {
	// a := [2]int{}//[0 0]
	// a := [2]int{1}      //[1 0]
	// a := [5]int{4: 1} //[0 0 0 0 1]
	// a := [...]int{1, 2, 3, 4, 5}//
	// a := [...]int{0: 1, 1: 2, 2: 3}
	a := [...]int{19: 1}
	var p *[20]int = &a //指向数组的指针
	fmt.Println(a)
	fmt.Println(p)

	x, y := 1, 2
	b := [...]*int{&x, &y} //指针数组
	fmt.Println(b)
	c := [2]int{1, 2}
	d := [2]int{1, 2}
	fmt.Println(c == d)

	e := new([10]int) //返回一个指向数组的指针
	fmt.Println(e)
}
