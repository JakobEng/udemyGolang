package main

import "fmt"

func main() {
	switch 2 {
	case 0:
		fmt.Println("This is the worst resturant ever")
	case 1:
		fmt.Println("This is just a bad resturant, don't come back")
	case 2:
		fmt.Println("This resturant did alright, I will come back again")
		fallthrough
	case int:
		fmt.Println("The best resturant every made")
	}
}
