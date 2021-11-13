package main

import "fmt"

func main_array() {
	var a [5]int
	fmt.Println("Empty array: ", a)
	fmt.Println("Array's length: ", len(a))
	a[0] = 57
	a[1] = 92
	a[2] = 78
	a[3] = 62
	a[4] = 18
	fmt.Println("initialized array: ", a)
	fmt.Println("get element at position 2: ", a[2])

	b := [5]int{3, 7, 98, 426, 156}
	fmt.Println("Declared and initialized ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Two Dimensional Array", twoD)
}
