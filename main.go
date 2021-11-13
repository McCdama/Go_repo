package main

import (
	"fmt"
	ts "gorepo/typestruct"
)

func main() {

	// artes.Main_array()
	// constant.Main_constant()
	// fory.Main_for()
	// hello.Main_hello()
	// ify.Main_if()
	// infer.Main_infer()
	// mulreturn.Main_multiple_return()
	// slicey.Main_slice()
	// switchy.Main_switch()

	billy := ts.ConBill("Bon_1", map[string]float64{"Chicken pot pie": 30.4, "Meatloaf": 12.8}, 1.55)
	fmt.Println(billy.Format()) // Receiver Function

}
