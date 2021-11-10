package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's befoor noon")
	default:
		fmt.Println("It's after noon")
	}

	/* A type switch compares types instead of values */

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("boolean")
		case int:
			fmt.Println("Integer")
		default:
			fmt.Printf("Not discovered yet", t)
		}
	}

	whatAmI(true)
	whatAmI(33)
	whatAmI("Hey there delilah!")
}
