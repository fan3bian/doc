package main

import (
	"fmt"
	// "time"
)

func main() {
	//无缓存：阻塞(同步)
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO")
		c <- true //存入
		close(c)
	}()
	// <-c //取出

	for v := range c {
		fmt.Println(v)
	}

	// time.Sleep(2 * time.Second)

	//单向存入

}

// func Go() {
// 	fmt.Println("GO GO GO")
// }
