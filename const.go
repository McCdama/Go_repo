package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main_() {
	fmt.Println(s)

	const n = 5000000

	const d = 9e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
