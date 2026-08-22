package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/expr-lang/expr"
)

func calculator() {
	expression := strings.Join(os.Args[1:], " ")
	result, err := expr.Eval(expression, nil)

	if err != nil {
		fmt.Println("Use only numbers.")
		return
	} else {
		fmt.Printf("%v\n", result)
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Use the calculator like this: golc \"1 + 1 * 2 - 4\"")
		fmt.Println("+ = Addition")
		fmt.Println("- = Subtraction")
		fmt.Println("/ = Division")
		fmt.Println("\\* = Multiplication")
		fmt.Println("")
		fmt.Println("Only use \"\" when the numerical expression has multiplication")
		return
	} else {
		calculator()
	}
}
