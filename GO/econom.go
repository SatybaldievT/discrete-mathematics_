package main

import (
	"fmt"
	"os"
)

var m map[string]bool

func econom(a []rune) {
	if len(a) == 1 {
		return
	} else {

		m[string(a)] = true
		val := make([]rune, 0)
		val2 := make([]rune, 0)
		k := 0
		i := 2
		if a[i] == rune(40) {

			for i = i; i < len(a); i++ {
				if a[i] == 40 {

					k++
				}
				if a[i] == 41 {

					k--
					if k == 0 {
						val = append(val, a[i])
						break
					}

				}

				val = append(val, a[i])
			}
		} else {
			val = append(val, a[i])
		}

		k = 0
		i++
		if a[i] == rune(40) {

			for i = i; i < len(a); i++ {
				if a[i] == 40 {

					k++
				}
				if a[i] == 41 {

					k--
					if k == 0 {
						val2 = append(val2, a[i])
						break
					}

				}

				val2 = append(val2, a[i])
			}
		} else {
			val2 = append(val2, a[i])
		}

		if !m[string(val)] {
			fmt.Println("val1", string(val))
			econom(val)
		}
		if !m[string(val2)] {
			fmt.Println("val2", string(val2))
			econom(val2)
		}
	}
}
func main() {
	m = make(map[string]bool)
	var a string
	fmt.Fscanln(os.Stdin, &a)
	econom([]rune(a))
	fmt.Println(len(m))

}
