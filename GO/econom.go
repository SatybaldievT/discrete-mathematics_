package main

import (
	"fmt"
	"os"
)

var m map[string]int

func polish(a []rune) int {
	if len(a) == 1 {
		return (int(a[0]) - 48)
	} else {
		op := a[1]
		val := make([]rune, 0)
		val2 := make([]rune, 0)

		i := 2
		if a[2] == rune(40) {

			for i = 2; i < len(a); i++ {
				if a[i] == 41 {
					val = append(val, a[i])
					break
				}
				val = append(val, a[i])
			}
		} else {
			val = append(val, a[i])
		}

		if a[i+1] == 40 {

			for i = i + 1; i < len(a); i++ {
				if a[i] == 41 {
					val2 = append(val2, a[i])
					break
				}
				val2 = append(val2, a[i])
			}
		} else {
			val2 = append(val2, a[i+1])
		}

		switch op {
		case 42:
			return (polish(val) * polish(val2))
		case 43:
			return (polish(val) + polish(val2))
		case 45:
			return (polish(val) * polish(val2))
		}
		return 1945
	}
}
func main() {
	m = make(map[string]int)
	var a string
	fmt.Fscanln(os.Stdin, &a)
	fmt.Println(len(m))
	fmt.Println(polish([]rune(a)))

}
