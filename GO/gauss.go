package main

import (
	"fmt"
	"os"
)

type pair struct {
	a, b int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
func gcd(a, b int) int {
	a, b = abs(a), abs(b)

	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func lcm(a, b int) int {
	a, b = abs(a), abs(b)

	return (a * b) / gcd(a, b)
}
func sum(a, b pair) pair {
	q := b
	g := gcd(q.b, a.b)
	q.a = a.a*(q.b/g) + q.a*(a.b/g)
	q.b = q.b * a.b / g
	return q.norm()
}
func (a pair) norm() pair {
	g := gcd(a.a, a.b)
	a.a /= g
	a.b /= g
	if a.a < 0 && a.b < 0 {
		a.a = -a.a
		a.b = -a.b
	}
	if a.a > 0 && a.b < 0 {
		a.a = -a.a
		a.b = -a.b
	}
	return a
}
func sum_pair(q []pair, b []int) pair {
	a := q[0]
	a.a *= b[0]
	for i := 1; i < len(q); i++ {

		a = sum(a, pair{q[i].a * b[i], q[i].b})
	}
	return a
}
func gauss(smt [][]int) []pair {
	if len(smt) == 1 {
		//выход из рекурсии
		l := len(smt[0]) - 1
		if l > 2 {

			return nil
		}

		if smt[0][l-1] == 0 {

			return nil
		}
		if smt[0][l] == 0 {
			return nil
		}

		return append(make([]pair, 0), pair{smt[0][1], smt[0][0]}.norm())

	} else {
		a := smt[0]
		i := 0
		for i = 0; i < len(smt)-1; i++ {

			if smt[i][0] != 0 {
				break
			}
		}
		if i == len(smt)-1 {
			return nil
		}
		smt[0] = smt[i]
		smt[i] = a
		for i = 1; i < len(smt); i++ {
			for j := 1; j < len(smt)+1; j++ {
				smt[i][j] = smt[i][j]*smt[0][0] - smt[0][j]*smt[i][0]
			}
			smt[i][0] = 0
		}

		a = make([]int, len(smt[0]))
		copy(a, smt[0])

		smt = smt[1:]
		for i = 0; i < len(smt); i++ {
			smt[i] = smt[i][1:]
		}
		q := gauss(smt)
		if q == nil {
			return nil
		}

		su := sum_pair(q, a[1:len(a)-1])

		su.a *= -1
		su = sum(su.norm(), pair{a[len(a)-1], 1})
		su = pair{su.a, su.b * a[0]}
		return append([]pair{su.norm()}, q...)
	}

}
func main() {
	var a int
	fmt.Fscanln(os.Stdin, &a)
	smt := make([][]int, a)

	for i := 0; i < a; i++ {
		smt[i] = make([]int, a+1)
		for j := 0; j < a+1; j++ {
			fmt.Fscan(os.Stdin, &smt[i][j])

		}
	}
	g := gauss(smt)
	if g == nil {
		fmt.Println("No solution")
		return
	}
	for i := 0; i < len(g); i++ {
		fmt.Printf("%d/%d\n", g[i].a, g[i].b)
	}

}
