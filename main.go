package main

import (
	"fmt"
)

func main() {
	fmt.Println("for")

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop for")
		break
	}

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
		// if break here, i=0, i++ is Unreachable code
		break
	}
}
