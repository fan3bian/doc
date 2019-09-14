package main

import (
	"fmt"
)

func main() {
	a := [...]int{3, 5, 2, 3, 7, 4}
	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] > a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}

		}
	}
	fmt.Println(a)
}
