package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int, 3)
	go test1(c1)
	for v := range c1 {
		fmt.Println(v)
	}

	fmt.Println("after close c1 do something")
	c2 := make(chan bool)
	go func() {
		time.AfterFunc(time.Second * 3, func() {
			close(c2)
		})
	}()
	_, close := <- c2
	if !close {
		fmt.Println("after c2 closed do something")
	}

	fmt.Println("end")
}

func test1(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}
