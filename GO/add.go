package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func add(a, b []int32, w int) []int32 {
	size := max(len(a), len(b))
	x := make([]int32, 0, size+1)
	p := int32(w)
	var q int32 = 0
	for i := 0; i <= size; i++ {

		if len(a) > i {
			q = q + a[i]
		}
		if len(b) > i {
			q = q + b[i]
		}
		if q == 0 && i == size {
			return x
		}

		x = append(x, q%p)
		q = q / p
	}

	return x
}
func main() {
	a := []int32{7, 7, 7, 7, 7}
	b := []int32{-7, -7, -7, -7, -7}

	fmt.Println(add(a, b, 8))
}
