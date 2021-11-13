package main

import (
	"fmt"

	"github.com/McCdama/Go_repo/artes"
	"github.com/McCdama/Go_repo/constant"
	"github.com/McCdama/Go_repo/fory"
	"github.com/McCdama/Go_repo/hello"
	"github.com/McCdama/Go_repo/ify"
	"github.com/McCdama/Go_repo/infer"
	"github.com/McCdama/Go_repo/mulreturn"
	"github.com/McCdama/Go_repo/slicey"
	"github.com/McCdama/Go_repo/switchy"
	ts "github.com/McCdama/Go_repo/typestruct"
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

	billy := ts.ConBill("Bon_1", map[string]float64{"Chicken pot pie": 30.4, "Meatloaf": 12.8}, 1.55)
	fmt.Println(billy)

}
