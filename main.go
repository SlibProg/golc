package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculator() {
	number, err1 := strconv.Atoi(os.Args[1])
	op := os.Args[2]
	number2, err2 := strconv.Atoi(os.Args[3])

	if err1 != nil || err2 != nil {
		fmt.Println("Use only numbers.")
		return
	}

	if op == "+" {
		fmt.Println(number + number2)
	} else if op == "-" {
		fmt.Println(number - number2)
	} else if strings.ToUpper(op) == "X" {
		fmt.Println(number * number2)
	} else if op == "/" {
		if number == 0 {
			fmt.Println("Don't even try")
		} else if number2 == 0 {
			fmt.Println("Don't even try")
		} else {
			fmt.Println(number / number2)
		}
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Use the calculator like this:")
		fmt.Println("golc 1 + 1")
		fmt.Println("")
		fmt.Println("+ = Addition")
		fmt.Println("- = Subtraction")
		fmt.Println("/ = Division")
		fmt.Println("x = Multiplication")
		return
	} else {
		calculator()
	}
}
