package main

import (
	"fmt"
	"strings"
)

// get the initial values
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	// split up the string into array
	names := strings.Split(s, " ") // return a slice

	// another slice for the initials
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1]) // use the range all the way up to 1
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	// a different return
	return initials[0], "_"
}

func main_multiple_return() {
	fn1, sn1 := getInitials("mohammad alrahbi")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("mccdama")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("ajassem")
	fmt.Println(fn3, sn3)
}
