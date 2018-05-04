package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan bool, 0)  // unbuffered channel
	c2 := make(chan bool)     // unbuffered channel

	go func() {
		c1 <- false
		time.Sleep(time.Second * 2)
		c2 <- true
	}()

	fmt.Println(<-c1)
	fmt.Println(<-c2)

	c3 := make(chan bool, 5)  // buffered channel

	go func() {
		for i := 0; i < 20; i++ {
			c3 <- i % 2 == 0
		}
		close(c3)
	}()

	for v := range c3 {
		fmt.Println(v)
	}
}