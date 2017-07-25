package main

import "fmt"

func main() {
	payments := make([]map[string]interface{}, 0, 25)

	payments = append(payments,
		pay("today", 100),
		pay("yesterday", 500),
		pay("yesterday", 2))

	delete(payments[1], "time")

	fmt.Println(payments)
}

func pay(time string, amount float64) map[string]interface{} {
	return map[string]interface{}{
		"time":   time,
		"amount": amount,
	}
}
