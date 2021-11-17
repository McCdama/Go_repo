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

	// Declare a string pointer array of three elements.
	var array_6_0 [3]*string
	fmt.Println("Array_6_0 PONTER", array_6_0)
	// Declare a second string pointer array of three elements.
	// Initialize the array with string pointers.
	array_6_1 := [3]*string{new(string), new(string), new(string)}
	fmt.Println("Array_6_1 PONTER", array_6_1)
	// Add colors to each element
	*array_6_1[0] = "Red"
	*array_6_1[1] = "Blue"
	*array_6_1[2] = "Green"
	// Copy the values from array2 into array1.
	array_6_0 = array_6_1
	fmt.Println("Array_6_0 PONTER after copying", array_6_0)

	fmt.Println("-----Two dimensional-----")

	// Declare a two dimensional integer array of four elements by two elements.
	var array_d_0 [4][2]int
	fmt.Println("Array_d_0", array_d_0)

	// Use an array literal to declare and initialize a two  dimensional integer array.
	array_d_1 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println("Array_d_1", array_d_1)

	// Declare and initialize index 1 and 3 of the outer array.
	array_d_2 := [4][2]int{1: {20, 21}, 3: {40, 41}}
	fmt.Println("Array_d_2", array_d_2)

	// Declare and initialize individual elements of the outer and inner array.
	array_d_3 := [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println("Array_d_3", array_d_3)

	/*
		Passing an array between functions can be an expensive operation in terms of memory
		and performance.
		When you pass variables between functions, they’re always passed by value.
		When your variable is an array, this means the entire array,
		regardless of its size, is copied and passed to the function.

		// Declare an array of 8 megabytes.
		var array [1e6]int
		// Pass the array to the function foo.
		foo(array)
		// Function foo accepts an array of one million integers.
		func foo(array [1e6]int) {...}

		Every time the function foo is called, eight megabytes of memory has to be allocated
		on the stack.
		Then the value of the array, all eight megabytes of memory, has to be
		copied into that allocation.
		Go can handle this copy operation, but there’s a better
		and more efficient way of doing this. You can PASS A POINTER to the array and only copy
		eight bytes, instead of eight megabytes of memory on the stack.
		// Allocate an array of 8 megabytes.
		var array [1e6]int
		// Pass the address of the array to the function foo.
		foo(&array)
		// Function foo accepts a pointer to an array of one million integers.
		func foo(array *[1e6]int) {...}

		This time the function foo TAKES A POINTER to an array of one million elements of type integer.
		The function call now PASSES THE ADDRESS of the array, which only requires
		eight bytes of memory to be allocated on the stack for the pointer variable.
		This operation is much more efficient with memory and could yield better performance.

	*/

}
