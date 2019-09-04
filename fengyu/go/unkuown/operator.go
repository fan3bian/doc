package main

import (
	"fmt"
)

/*
6 ：0110
11：1011
--------
& : 0010
| : 1111
^ : 1101 //两个中有一个1一个0时成立
&^：0100 //第二个数取反
*/
func main() {
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)
	a := 1
	if a > 0 && (10/a) > 1 {
		fmt.Println("OK")
	}

	var p *int = &a
	fmt.Println(p)
	fmt.Println(*p)
}
