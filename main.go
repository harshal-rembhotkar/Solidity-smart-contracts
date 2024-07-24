package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var num1 int

	var num2 int

	var operator rune
	for {
		fmt.Print("enter the first number: ")
		fmt.Scanln(&num1)

		fmt.Print("enter the second number: ")
		fmt.Scanln(&num2)

		fmt.Print("Enter the operator:")
		fmt.Scanf("%c \n", &operator)

		if operator == '+' || operator == '-' || operator == '*' || operator == '/' || operator == '%' {

			if operator == '+' {
				fmt.Print("the result is :")
				fmt.Println(Addition(num1, num2))
			}

			if operator == '-' {
				fmt.Print("the result is :")
				fmt.Println(Subtraction(num1, num2))
			}

			if operator == '*' {
				fmt.Print("the result is :")
				fmt.Println(Multiplication(num1, num2))
			}

			if operator == '/' {
				fmt.Print("the result is :")
				fmt.Println(Division(num1, num2))
			}

			if operator == '%' {
				fmt.Print("the result is :")
				fmt.Println(Modular(num1, num2))
			}

		} else if operator == 'x' || operator == 'X' {
			fmt.Println("exited")
			return

		} else {

			fmt.Println("Invalid operator")
		}
	}

}
