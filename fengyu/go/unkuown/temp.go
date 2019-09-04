package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	e = 1
	f
	g
	h    = "sss"
	i    = len(h)
	j, k = 1, "2"
	l, m
)
const (
	Monday = iota
	Tuesday
)

type (
	//可以自定义,但不要使用
	byte int8
)

func main() {
	var a int
	// var b bool
	// var s string
	// var sli []int
	// var arr [2]int
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(s)
	// fmt.Println(sli)
	// fmt.Println(arr)

	fmt.Println(math.MinInt8)

	a = 65
	as := string(a)
	ass := strconv.Itoa(a)
	a, _ = strconv.Atoi(ass)
	fmt.Println(as)
	fmt.Println(ass)
	fmt.Println(a)
	fmt.Println(e, f, g)    //常量
	fmt.Println(j, k, l, m) //常量

}
