package main

import (
	"fmt"

	"github.com/Go_repo/artes"
	"github.com/Go_repo/constant"
	"github.com/Go_repo/fory"
	"github.com/Go_repo/hello"
	"github.com/Go_repo/ify"
	"github.com/Go_repo/infer"
	"github.com/Go_repo/mulreturn"
	"github.com/Go_repo/slicey"
	"github.com/Go_repo/switchy"
	ts "github.com/Go_repo/typestruct"
)

func main() {
	artes.Main_array()
	constant.Main_constant()
	fory.Main_for()
	hello.Main_hello()
	ify.Main_if()
	infer.Main_infer()
	mulreturn.Main_multiple_return()
	slicey.Main_slice()
	switchy.Main_switch()

	// typestruct package
	billy := ts.ConBill("Bon_1", map[string]float64{"Chicken pot pie": 30.4, "Meatloaf": 12.8}, 1.55)
	fmt.Println(billy)

}
