package main

import "fmt"

func main() {
	fmt.Println(average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func average(nums ...float64) float64 {
	sum := 0.0
	for _, v := range nums {
		sum += v
	}
	return sum / float64(len(nums))
}
