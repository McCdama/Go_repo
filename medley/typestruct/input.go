package typestruct

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func PromptOpt(b Bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := GetInput("Choose option (a - add item, s - save bill, t - add tip, q - for exit):", reader)
	// fmt.Print(opt)

	switch opt {
	case "a":
		name, _ := GetInput("Item name: ", reader)
		price, _ := GetInput("Item price: ", reader)

		// parsing string into float
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			PromptOpt(b) // go back
		}
		b.AddItem(name, p)
		fmt.Println("Item added - ", name, price)
		PromptOpt(b) // add another item, tip, or save
	case "s":
		b.Save()
		fmt.Println("you saved the file - ", b.Name)
	case "t":
		tip, _ := GetInput("Tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			PromptOpt(b) // go back
		}
		b.UpdateTip(t)

		fmt.Println("tip added - ", tip)
		PromptOpt(b)

	case "q":
		break
	default:
		fmt.Println("Please enter a valid letter... ")
		PromptOpt(b) // recycle the func
	}
}
