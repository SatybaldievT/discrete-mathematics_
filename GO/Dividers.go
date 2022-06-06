package main

import (
	"fmt"
	"math"
	"os"
)

func dev(a int, b []int) {
	for _, q := range b {
		if a%q == 0 {
			fmt.Printf("%d--%d;\n", a, a/q)
			dev(a/q, b)
		}
	}
}
func main() {
	var a int
	fmt.Fscanln(os.Stdin, &a)
	seive := make([]bool, int(math.Sqrt(float64(a)))+1)

	limit := int(math.Sqrt(math.Sqrt(float64(a)))) + 1

	for i := 2; i < limit; i++ {
		if !seive[i] {
			for j := i * i; j < int(math.Sqrt(float64(a))); j += i {
				seive[j] = true
			}
		}
	}
	p := make([]int, 0)
	fmt.Println("graph {")
	for i := 2; i < int(math.Sqrt(float64(a))); i++ {
		if !seive[i] {
			p = append(p, i)
		}
	}
	dev(a, p)
	fmt.Println("}")
}
