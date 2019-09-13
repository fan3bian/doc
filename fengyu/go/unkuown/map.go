package main

import (
	"fmt"
	"sort"
)

func main() {
	var m map[int]string
	m = map[int]string{}
	m = make(map[int]string)
	// var n map[int]string = make(map[int]string)
	n := make(map[int]string)
	n[1] = "Ok"
	fmt.Println(m)
	fmt.Println(n)
	delete(n, 1)
	fmt.Println(n)

	k := make(map[int]map[int]string)
	k[1] = make(map[int]string)
	k[1][1] = "OK"
	// a := k[1][1]
	a, exists := k[2][1]
	if !exists {
		k[2] = make(map[int]string)
	}
	k[2][1] = "good"
	a = k[2][1]
	fmt.Println(a, exists)

	//iterate
	sm := make([]map[int]string, 5)
	for _, v := range sm {
		v = make(map[int]string, 1)
		v[1] = "ok"
		fmt.Println(v)

	}
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "ok"

	}
	fmt.Println(sm)
	//map进行遍历
	tm := map[int]string{1: "a", 2: "b", 3: "c"}
	s := make([]int, len(tm))
	i := 0
	for k, _ := range tm {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)
}
