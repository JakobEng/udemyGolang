/*

The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.

*/
package main

import "fmt"

func main() {
	var sum uint

	for i := 1; i <= 1000; i++ {
		var thisPower uint = uint(i)

		for a := 1; a <= i-1; a++ {
			thisPower *= uint(i)
		}

		sum += thisPower
	}
	fmt.Println(sum)
}
