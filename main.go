package main

import (
	"fmt"
)

func main() {
	fmt.Println("arrays")
	var a [5]int
	fmt.Println("zero arrays: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len(a) ", len(a))

	b := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("dcl: ", b)

	var towD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			towD[i][j] = i + j
		}
	}

	fmt.Println("tow: ", towD)
}
