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

	opt, _ := GetInput("Choose option (a - add item, s - save bill, t - add tip):", reader)
	fmt.Print(opt)
}
