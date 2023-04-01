package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("switch")
	i := 2
	fmt.Printf("%d as english ", i)
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("it is weekend")
	default:
		fmt.Println("is work day")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("is up time")
	default:
		fmt.Println("default")
	}

	whoAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("is BOOL type")
		case int:
			fmt.Println("is INT type")
		case string:
			fmt.Println("is STRING type")
		default:
			fmt.Printf("un know type: %T\n", t)
		}
	}

	whoAmI(true)
	whoAmI(2)
	whoAmI("string")
	whoAmI(struct{}{})

}
