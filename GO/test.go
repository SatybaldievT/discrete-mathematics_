package main

import (
	"fmt"
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
	return a
}

/*func sum_pair(a []pair, b []int) pair {

	a[0].a *= b[0]
	for i := 1; i < len(a); i++ {
		a[i].a *= b[i]
		a[0] = sum(a[0], a[i])
	}
	return a[0]
}*/
func sum_pair(q []pair, b []int) pair {
	a := q[0]
	a.a *= b[0]
	for i := 1; i < len(q); i++ {

		a = sum(a, pair{q[i].a * b[i], q[i].b})
	}
	return a
}
func main() {
	q := []pair{pair{214, 7}, pair{274, 21}}
	a := []int{-4, -1, 8, 2}
	fmt.Println(q, a, sum_pair(q, a[1:len(a)-1]))
}
