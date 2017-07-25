package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	for i := 0; i <= 100; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Cap:", cap(mySlice), "Val:", mySlice[len(mySlice)-1])
	}
}
