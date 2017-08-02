package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		time.Sleep(time.Second)
		close(c)
	}()

	fmt.Println("Wating")
	fmt.Println(<-c)
	fmt.Println("Done")

}
