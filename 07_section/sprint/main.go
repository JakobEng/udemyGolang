package main

import (
	"fmt"
	"strconv"
)

func main() {
	productID := 42
	fmt.Println("Product with ID '" + strconv.Itoa(productID) + "' could not be found.")
}
