package main

import (
	"fmt"
)

func main() {

	var val1, val2, result float64
	var operator string

	fmt.Print("Please enter first number: ")
	fmt.Scan(&val1)

	fmt.Print("Please enter second number: ")
	fmt.Scan(&val2)

	fmt.Print("Choose your operator: (+, -, *, /)")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		result = val1 + val2
	case "-":
		result = val1 - val2
	case "*":
		result = val1 * val2
	case "/":
		if val2 != 0 {
			result = val1 / val2
		} else {
			fmt.Println("Division by zero is not allowed")
			return
		}
	default:
		fmt.Println("Invalid operator")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", val1, operator, val2, result)

}
