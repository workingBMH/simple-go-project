package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Addition: ", Add(2, 3))
	fmt.Println("Subtraction: ", Subtract(5, 2))
	fmt.Println("Multiplication: ", Multiply(3, 4))
	fmt.Println("Division: ", Divide(10, 2))
	fmt.Println("Uppercase: ", ToUpperCase("hello"))
}

// Add returns the sum of a and b
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of a and b
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of a and b
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of a and b
func Divide(a, b int) int {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}

// ToUpperCase converts a string to uppercase
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}
