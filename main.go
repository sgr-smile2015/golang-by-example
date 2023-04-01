package main

import (
	"fmt"
)

func main() {
	fmt.Println("slices")

	s := make([]string, 3)
	fmt.Println("initial slice: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len:  ", len(s))
	s = append(s, "e")
	s = append(s, "f", "g")
	fmt.Println("append slice: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy s slice: ", c)

	l := s[2:5]
	fmt.Println("slice [2:5]", l)
	l = s[:5]
	fmt.Println("slice [:5]", l)
	l = s[2:]
	fmt.Println("slice [2:]", l)

	t := []string{"java", "golang", "python", "c++", "rust"}
	fmt.Println("t: ", t)

	towD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		towD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			towD[i][j] = i + j
		}
	}

	fmt.Println("towD: ", towD)

}
