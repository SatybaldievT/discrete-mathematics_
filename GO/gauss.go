package main

import (
	"fmt"
	"os"
)

func gauss(smt [][]int) {
	if len(smt) == 1 {
		//выход из рекурсии
		/*l:=len(smt[0])
		for i:=0;i<l-1;i++{
			if smt[0][i] != 0 {return make([]pair,0),false}
		}
		if smt[0][l-1]==0 && smt[0][l]!=0 {return make([]pair,0),false}
		if smt[0][l-1]==0 && smt[0][l]==0 {return append(make([]pair, 0), pair{0,1}),true}
		return append(make([]pair, 0), pair{smt[0][l-1],smt[0][l]}),true*/
		l := len(smt[0]) - 1
		if l > 2 {
			fmt.Println("ERROR")
			return
		}

		if smt[0][l-1] == 0 {
			fmt.Println("ERROR")
			return
		}
		if smt[0][l] == 0 {
			fmt.Println("ERROR")
			return
		}

		fmt.Print(smt[0])

	} else {
		a := smt[0]
		i := 0
		for i = 0; i < len(smt)-1; i++ {

			if smt[i][0] != 0 {
				break
			}
		}
		if i == len(smt)-1 {
			fmt.Println("ERROR")
			return
		}
		smt[0] = smt[i]
		smt[i] = a
		for i = 1; i < len(smt); i++ {
			for j := 1; j < len(smt)+1; j++ {
				smt[i][j] = smt[i][j]*smt[0][0] - smt[0][j]*smt[i][0]
			}
			smt[i][0] = 0
		}
		fmt.Println(smt[0])
		smt = smt[1:]
		for i = 0; i < len(smt); i++ {
			smt[i] = smt[i][1:]
		}
		gauss(smt)
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
	gauss(smt)

}
