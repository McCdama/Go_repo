package typestruct

import (
	"bufio"
	"fmt"
	"os"
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
		price, _ := GetInput("Item price: ", reader) // we got string from input!--> turn into float
		fmt.Println(name, price)

	case "s":
		fmt.Println("you chose s")
	case "t":
		tip, _ := GetInput("Tip amount ($): ", reader)
		fmt.Println(tip)

	case "q":
		break
	default:
		fmt.Println("Please enter a valid letter... ")
		PromptOpt(b) // recycle the func
	}
}
