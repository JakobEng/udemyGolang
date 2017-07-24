package main

import "fmt"

func main() {
	for i := 10000; i < 10100; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), 'a')
	}
}
