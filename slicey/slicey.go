package slicey

import "fmt"

func Main_slice() {
	/*
	* Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	* To create an empty slice with non-zero length, --> make
	* Here we make a slice of strings of length 3 (initially zero-valued).
	 */
	s := make([]string, 3)
	fmt.Println(s)
	s[0] = "Lorem ipsum dolor sit amet"
	s[1] = "onsectetur adipiscing elit,"
	s[2] = "ed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	fmt.Println("Initialized values are: ", s)
	fmt.Println("-------------------------")
	fmt.Println("Get a specifed value on position 2:", s[2])
	fmt.Println("-------------------------")
	fmt.Println("Expected length", len(s))
	fmt.Println("-------------------------")

	s = append(s, "Imperdiet massa tincidunt nunc pulvinar sapien et ligula.")
	fmt.Println("New length", len(s))
	fmt.Println("After appending", s)
	fmt.Println("-------------------------")

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy'd Slice", c)
	fmt.Println("-------------------------")

	// Slice operator [low:high]
	l := s[2:5]
	fmt.Println("Sliced slice", l)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Print("Two Dimensional Slice: ", twoD)

	// further info, checkout: https://go.dev/blog/slices-intro
}
