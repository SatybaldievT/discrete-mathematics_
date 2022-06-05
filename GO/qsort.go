package main

import "fmt"

func Partition(low int, high int, less func(i, j int) bool,
	swap func(i, j int)) int {
	i := low
	j := low
	for j < high {
		if less(j, high) {
			swap(i, j)
			i = i + 1
		}
		j = j + 1

	}
	swap(i, high)
	return i
}
func QuickSortRec(low int, high int, less func(i, j int) bool,
	swap func(i, j int)) {
	if low < high {
		q := Partition(low, high, less, swap)
		QuickSortRec(low, q-1, less, swap)
		QuickSortRec(q+1, high, less, swap)
	}
}
func qsort(n int,
	less func(i, j int) bool,
	swap func(i, j int)) {
	QuickSortRec(0, n-1, less, swap)
}

var a = []int{-1, 8, 7, 6, 5, 4, 3, 2, 1}

func less(q int, b int) bool {
	return a[q] < a[b]
}
func swap(q int, b int) {
	w := a[q]
	a[q] = a[b]
	a[b] = w
}
func main() {

	qsort(9, less, swap)
	for i := 0; i < 9; i++ {
		fmt.Println(a[i])
	}
}
