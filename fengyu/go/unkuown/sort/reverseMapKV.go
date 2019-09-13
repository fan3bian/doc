package main

import (
	"fmt"
)

func main() {
	m := map[int]string{1: "a", 2: "b"}
	n := make(map[string]int)

	for k, v := range m {
		// t := m[k]
		// fmt.Println(t)
		n[v] = k
	}
	fmt.Println(m, n)
}
