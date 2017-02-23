package main

import (
	"fmt"
	"os"
	"strconv"
)

func multiply(numbers []string) {
	result := 1

	for _, param := range numbers {
		numberToMultiply, _ := strconv.Atoi(param)
		result = numberToMultiply * result
	}
	fmt.Println(result)
}

func add(numbers []string) {
	result := 0

	for _, param := range numbers {
		numberToAdd, _ := strconv.Atoi(param)
		result = numberToAdd + result
	}
	fmt.Println(result)
}

func divide(numbers []string) {
	result, _ := strconv.Atoi(numbers[0])

	numbers = append(numbers[:0], numbers[1:]...)

	hasError := false

	for _, param := range numbers {
		numberToDivide, _ := strconv.Atoi(param)
		if numberToDivide == 0 {
			fmt.Println("dividing by 0 is impossible, sir.")
			hasError = true
			break
		} else {
			result = result / numberToDivide
		}
	}
	if !hasError {
		fmt.Println(result)
	}
}

func substract(numbers []string) {
	result, _ := strconv.Atoi(numbers[0])

	numbers = append(numbers[:0], numbers[1:]...)

	for _, param := range numbers {
		numberToSubstract, _ := strconv.Atoi(param)
		result = result - numberToSubstract
	}
	fmt.Println(result)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("You must enter at least three parameters : an operand and two numbers.")
	} else {
		operator := os.Args[1]
		params := os.Args[2:]
		switch operator {
		case "+":
			add(params)

		case "-":
			substract(params)

		case "*":
			multiply(params)

		case "/":
			divide(params)

		default:
			fmt.Println("You entered an invalid operator. Please enter one of the following : +, -,\"*\",/")
		}
	}
}
