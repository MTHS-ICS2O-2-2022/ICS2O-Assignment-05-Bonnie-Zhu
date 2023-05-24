package main

import (
	"fmt"
	"strings"
)

func generatePyramid(inputNumber int) string {
	if inputNumber < 1 {
		return "Please enter a positive number."
	}

	var pyramid strings.Builder
	for i := 1; i <= inputNumber; i++ {
		spaces := strings.Repeat(" ", inputNumber-i)
		asterisks := strings.Repeat("*", i*2-1)
		pyramid.WriteString(spaces + asterisks + "\n")
	}

	return pyramid.String()
}

func main() {
	var inputNumber int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scanln(&inputNumber)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	result := generatePyramid(inputNumber)
	fmt.Println(result)
}
