package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("Enter which operation? (add, subtract, multiply, divide)")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter first number")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter second number")
	fmt.Scanf("%d", &number2)

	switch operation {
	case "add":
		fmt.Println(number1 + number2)

	case "subtract":
		fmt.Println(number1 - number2)

	case "multiply":
		fmt.Println(number1 * number2)

	case "divide":
		fmt.Println(number1 / number2)
	}

}
