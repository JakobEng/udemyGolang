package main

import "fmt"

type square struct {
	side float64
}

type shape interface {
	area() float64
}

func (z square) area() float64 {
	return z.side * z.side
}

func info(x shape) {
	fmt.Println(x.area())
}

func main() {
	g := square{10}
	info(g)
}
