package main

import (
	"fmt"
)

func main() {
	var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice []int
	fmt.Println(array)
	fmt.Println(slice)
	s1 := array[5:10]
	s2 := array[5:]
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := make([]int, 3, 6)
	s4 := make([]int, 3)
	fmt.Println(len(s3), cap(s3))
	fmt.Println(len(s4), cap(s4))

	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'}
	sa := a[2:5]
	fmt.Println(string(sa))
	sb := sa[3:5]
	fmt.Println(string(sb))
	fmt.Printf("%v %p\n", s3, s3) //print addr
	s3 = append(s3, 1, 2, 3)
	fmt.Printf("%v %p\n", s3, s3)
	s3 = append(s3, 1, 2, 3)
	fmt.Printf("%v %p\n", s3, s3) //append enhance the cap automately

	arr := []int{1, 2, 3, 4, 5}
	s5 := arr[2:5]
	s6 := arr[1:3]
	fmt.Println(s5, s6)

	//copy
	s7 := []int{1, 2, 3, 4, 5, 6}
	s8 := []int{7, 8, 9}
	// copy(s7, s8)//target,source
	// fmt.Println(s7, s8)
	copy(s8, s7)
	fmt.Println(s7, s8)
	s9 := s7[:]
	fmt.Printf("%p %p", s7, s9)

}
