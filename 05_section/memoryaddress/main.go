package main

import "fmt"

func main() {
	var km float64
	fmt.Print("how long did you run:")
	fmt.Scan(&km)
	miles := km * 1.62
	fmt.Println(km, " km run or ", miles, "miles run.")
}
