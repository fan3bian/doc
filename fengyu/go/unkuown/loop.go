package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println(a)
	/*	for {
			a++
			if a > 3 {
				break
			}
		}
		fmt.Println(a)*/

	/*switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")

	}*/
	// switch {
	switch a := 1; {
	case a >= 0:
		fmt.Println("a=0")
		fallthrough
	case a >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}

}
