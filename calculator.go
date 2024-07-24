package main

import "fmt"

func Addition(num1 int, num2 int) int {
	return num1 + num2
}

func Subtraction(num1, num2 int) int {
	return num1 - num2
}

func Multiplication(num1 int, num2 int) int {

	return num1 * num2
}

func Division(num1 int, num2 int) float32 {
	var result int
	if num2 == 0 {
		fmt.Println("Cannot divide by zero")
		return 0
	}
	result = (num1 / num2)
	return float32(result)
}

func Modular(num1, num2 int) int {
	return num1 % num2
}
