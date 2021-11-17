package main

import "fmt"

func main() {
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

	array_1 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array_1", array_1)

	array_2 := [...]int{10, 20, 30, 40, 50}
	fmt.Println("Array_2", array_2)

	array_3 := [5]int{1: 10, 2: 20}
	fmt.Println("Array_3", array_3)

	array_2[2] = 35
	fmt.Println("Array_2", array_2)

	// Declare an integer pointer array of five elements.
	// Initialize index 0 and 1 of the array with integer pointers.
	array_4 := [5]*int{0: new(int), 1: new(int)}
	// Assign values to index 0 and 1.
	*array_4[0] = 10
	*array_4[1] = 20
	fmt.Println("Array_pointer", array_4)

	// Declare a string array of five elements.
	var array_5_0 [5]string
	fmt.Println("Array_5_0", array_5_0)
	// Declare a second string array of five elements.
	// Initialize the array with colors.
	array_5_1 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println("Array_5_1", array_5_1)
	// Copy the values from array2 into array1.
	array_5_0 = array_5_1
	fmt.Println("Array_5_0", array_5_0)
}
