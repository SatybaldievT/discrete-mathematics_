package main

import (
	"fmt"
	"os"
	"sort"
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
	var n int
	fmt.Fscanln(os.Stdin, &n)
	q := n
	a := []int{}
	p := []int{}
	d := 2
	for d*d <= n {
		if n%d == 0 {
			if len(p) != 0 && p[len(p)-1] == d {
				n = n / d
				continue
			}
			p = append(p, d)
			n = n / d
		} else {
			d += 1
		}
	}

	if n > 1 {
		if len(p) != 0 && p[len(p)-1] != n {
			p = append(p, n)
		}
		if len(p) == 0 {
			p = append(p, n)
		}
	}
	for d = 1; d*d < q; d++ {
		if q%d == 0 {
			a = append(a, d)
			a = append(a, q/d)
		}
	}
	if d*d == q {
		a = append(a, d)
	}
	sort.SliceStable(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	fmt.Println("graph {")
	for c, i := range a {
		if c != 0 && a[c-1] == i {
			continue
		}
		fmt.Printf("\t%d\n", i)
	}
	for c, i := range a {
		if c != 0 && a[c-1] == i {
			continue
		}
		for _, l := range p {
			if i%l == 0 {
				fmt.Printf("\t%d--%d\n", i, i/l)
			}
		}
	}

	fmt.Println("}")
}
