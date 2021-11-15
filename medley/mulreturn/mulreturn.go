package mulreturn

import (
	"fmt"
	"strings"
)

// get the initial values
func GetInitials(n string) (string, string) {
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

func Main_multiple_return() {
	fn1, sn1 := GetInitials("mohammad alrahbi")
	fmt.Println(fn1, sn1)

	fn2, sn2 := GetInitials("mccdama")
	fmt.Println(fn2, sn2)

	fn3, sn3 := GetInitials("ajassem")
	fmt.Println(fn3, sn3)
}
