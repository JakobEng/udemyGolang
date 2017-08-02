package main

import "fmt"

func main() {
	sums := factori(start(3, 13))

	for n := range sums {
		fmt.Println(n)
	}
}

func start(min, max int) <-chan int {
	out := make(chan int)
	go func(min, max int) {
		for i := min; i <= max; i++ {
			out <- i
		}
		close(out)
	}(min, max)
	return out
}

func factori(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
