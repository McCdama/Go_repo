package typestruct

import (
	"bufio"
	"fmt"
	"os"
)

type Bill struct {
	Name  string
	Items map[string]float64
	Tip   float64
}

// make a new Bill --> Like Constructor
func ConBill(n string, i map[string]float64, t float64) Bill {
	b := Bill{
		Name:  n,
		Items: i,
		Tip:   t,
	}
	return b
}

func (b Bill) Format() string {
	fs := "BILL INFORMATION: \n \n"
	var total float64 = 0

	for k, v := range b.Items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-25v ...$%0.2f", "\n total:", total)
	return fs
}

func CreateBill() Bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := GetInput("Create a new bill name: ", reader)

	b := ConBill(name, map[string]float64{}, 0)
	fmt.Println("Created the bill", b.Name)

	return b
}
