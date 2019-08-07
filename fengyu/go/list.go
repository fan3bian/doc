package main

import "fmt"
import "container/list"

func main() {
	l := list.New()
	l.PushBack("fist")
	l.PushFront(67)
	fmt.Println(l)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}