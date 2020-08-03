package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(basics("2+2*3+2"))
}

func slicerNumbers(expression string) []float64 {
	match := strings.FieldsFunc(expression, split)
	var result = []float64{}

	for _, i := range match {
		j, err := strconv.ParseFloat(i, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, j)
	}
	return result
}
func split(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
}
func slicerSymbols(expression string) []string {
	return strings.FieldsFunc(expression, splitNeg)
}
func splitNeg(r rune) bool {
	return r != '+' && r != '-' && r != '*' && r != '/'
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func removeNumber(slice []float64, s int, t rune) []float64 {
	if (s + 1) < len(slice) {
		if t == '*' {
			slice[s] *= slice[s+1]
		}
		if t == '/' {
			slice[s] /= slice[s+1]
		}

		if t == '+' {
			slice[s] += slice[s+1]
		}
		if t == '-' {
			slice[s] -= slice[s+1]
		}

		return append(slice[:s+1], slice[s+2:]...)
	}
	return []float64{slice[0]}

}

func removeSymbols(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)

}

func basics(expression string) float64 {
	numbers := slicerNumbers(expression)
	symbols := slicerSymbols(expression)
	count := len(symbols)

	for i := 0; i < count && len(symbols) != 0; i++ {
		switch symbols[i] {
		case "*":
			numbers = removeNumber(numbers, i, '*')
			symbols = removeSymbols(symbols, i)
			count -= 1
			i -= 1
		case "/":
			numbers = removeNumber(numbers, i, '/')
			symbols = removeSymbols(symbols, i)
			count -= 1
			i -= 1
		}
	}

	count = len(symbols)

	for i := 0; len(symbols) != 0; i++ {
		switch symbols[i] {
		case "+":
			numbers = removeNumber(numbers, i, '+')
			symbols = removeSymbols(symbols, i)
			count -= 1
			i -= 1
		case "-":
			numbers = removeNumber(numbers, i, '-')
			symbols = removeSymbols(symbols, i)
			count -= 1
			i -= 1
		}
	}
	return numbers[0]
}
