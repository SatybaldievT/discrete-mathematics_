package main

import "fmt"

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func main() {
	fmt.Println(gcd(5, 15))
}
