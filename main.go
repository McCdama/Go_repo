package main

import con "gorepo/concurrency"

// import ts "gorepo/typestruct"

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

	// billy := ts.NewBill("Bon_1", map[string]float64{"Chicken pot pie": 30.4, "Meatloaf": 12.8}, 1.55)
	// fmt.Println(billy.Format()) // Receiver Function

	// recieverBill := ts.NewBill("Bon_1", map[string]float64{"Chicken pot pie": 70.1, "Meatloaf": 11.29}, 3.94)
	// recieverBill.AddItem("Oninon Soup", 24.76)
	// fmt.Println(recieverBill.Format())

	// mybill := ts.CreateBill()
	// ts.PromptOpt(mybill)

	/*
		shapes := []i.Shape{
			i.Square{Length: 15.2},
			i.Circle{Radius: 7.5},
			i.Circle{Radius: 12.3},
			i.Square{Length: 4.9},
		}

		for _, v := range shapes {
			i.PrintShapeInfo(v)
			fmt.Println("---")
		}
	*/
	con.Run()
}
