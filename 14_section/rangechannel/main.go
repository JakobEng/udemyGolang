package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println(c)
		}
		fmt.Println("done")
		close(c)
	}()

	time.Sleep(time.Second)
	for n := range c {
		fmt.Println(n)
	}
}
