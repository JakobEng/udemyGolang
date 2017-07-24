package main

import "fmt"

type test string

func main() {
	var hello test = "hello"
	funcCall(hello)
}

func funcCall(z string) {
	fmt.Println(z)
}
