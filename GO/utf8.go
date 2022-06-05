package main

import "fmt"

func encode(utf32 []byte) []rune {
	return []rune(string(utf32))
}
func decode(utf32 []rune) []byte {
	return []byte(string(utf32))
}
func main() {
	rs := []rune{'H', 'e', 'l', 'l', 'o', ' ', '世', '界'}
	fmt.Println(rs)
	fmt.Println(decode(rs))
}
