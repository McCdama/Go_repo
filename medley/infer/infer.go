package infer

import (
	"fmt"
)

func Main_infer() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	/* The := syntax is shorthand for declaring and initializing a variable,
	e.g. for var f string = "Go" in this case. */
	f := "Go"
	fmt.Println(f)
}
