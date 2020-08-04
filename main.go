package main

import (
	"fmt"
)

func main() {
	expression := "(1,3,x)*2"
	fmt.Print(expression + " = ")
	displayVector(basicsVectors(expression))
}

func basics(expression string) float64 {

	for checkParenthesis(expression) {
		expression = handleParenthesis(expression)
	}

	numbers := slicerNumbers(expression)
	symbols := slicerSymbols(expression)
	count := len(symbols)

	for i := 0; i < count && len(symbols) != 0; i++ {
		switch symbols[i] {
		case "*":
			numbers = removeNumber(numbers, i, '*')
			symbols = removeSymbols(symbols, i)
			count--
			i--
		case "/":
			numbers = removeNumber(numbers, i, '/')
			symbols = removeSymbols(symbols, i)
			count--
			i--
		}
	}

	count = len(symbols)

	for i := 0; len(symbols) != 0; i++ {
		switch symbols[i] {
		case "+":
			numbers = removeNumber(numbers, i, '+')
			symbols = removeSymbols(symbols, i)
			count--
			i--
		case "-":
			numbers = removeNumber(numbers, i, '-')
			symbols = removeSymbols(symbols, i)
			count--
			i--
		}
	}
	return numbers[0]
}

func basicsVectors(expression string) []string {
	vectors := [][]string{}
	for checkParenthesis(expression) {
		vector, e := handleParenthesisVectors(expression)
		expression = e
		vectors = append(vectors, vector)
	}
	symbols := slicerSymbols(expression)
	numbers := slicerNumbers(expression)

	count := len(symbols)

	for i := 0; len(symbols) != 0; i++ {
		switch symbols[i] {
		case "*":
			vectors = removeVector(vectors, i, '*', numbers)
			symbols = removeSymbols(symbols, i)
			count--
			i--
		}
	}

	for i := 0; len(symbols) != 0; i++ {
		switch symbols[i] {
		case "+":
			vectors = removeVector(vectors, i, '+', numbers)
			symbols = removeSymbols(symbols, i)
			count--
			i--
		case "-":
			vectors = removeVector(vectors, i, '-', numbers)
			symbols = removeSymbols(symbols, i)
			count--
			i--
		}
	}
	return vectors[0]
}
