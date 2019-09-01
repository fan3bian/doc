package main

import(
	"fmt"
	"math"
 )
type (
	//可以自定义,但不要使用
	byte int8
)

func main() {
	var a int
	var b bool
	var s string
	var sli []int
	var arr [2]int
	fmt.Println(a)	
	fmt.Println(b)	
	fmt.Println(s)	
	fmt.Println(sli)	
	fmt.Println(arr)	
	
	fmt.Println(math.MinInt8)
}
