package main

import (
	"fmt"
	"strconv"
)

func main() {
	expression := "2+3+(1+5)+4+(1+3+(2+3))+(3)+(2+(2))"
	fmt.Println(expression + " = " + strconv.FormatFloat(basics(expression), 'f', -1, 64))
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
