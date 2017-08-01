package main

import (
	"fmt"
	"sort"
)

type people []int

func main() {
	group := people{2, 4, 3, 5, 1}
	sort.Ints(group)
	fmt.Println(group)
}
