package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("golfram>")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		parse := strings.Fields(text)

		if len(parse) == 0 {
			fmt.Println("Please parse an expression")
		} else if len(parse) == 1 {
			if parse[0] == "help" || parse[0] == "h" {
				fmt.Println("Welcome to Golfram")
				fmt.Println("basic, b, and an expression to solve it. Example : \" basic 2+3/4*5 \"")
				fmt.Println("vector, v, to work on vector. Example : \" vector (2,y)*z+(x,6)*2 \"")
				fmt.Println("degree, d, to get the degree of an expression. Example : \" degree x->x^3+x^2+x+1 \"")
				fmt.Println("function, f, to get the result of the function with a parameter. Example : \" function x->x+2 2 \"")
			} else {
				fmt.Println("Please parse an expression")
			}
		} else if len(parse) == 2 {
			if parse[0] == "basic" || parse[0] == "b" {
				fmt.Println(basics(parse[1]))
			} else if parse[0] == "vector" || parse[0] == "v" {
				displayVector(basicsVectors(parse[1]))
			} else if parse[0] == "degree" || parse[0] == "d" {
				fmt.Println(getDegree(parse[1]))
			} else {
				fmt.Println("Please parse a correct argument")
			}

		} else if len(parse) == 3 {
			if parse[0] == "function" || parse[0] == "f" {
				fmt.Println(simplifyFunc(parse[1], parse[2]))
			} else {
				fmt.Println("Please parse correct arguments")
			}

		} else {
			fmt.Println("Please parse an expression (2 arguments max)")
		}
	}
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
		case "^":
			numbers = removeNumber(numbers, i, '^')
			symbols = removeSymbols(symbols, i)
			count--
			i--
		}
	}

	count = len(symbols)

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
	numbers := slicerNumbersLetters(expression)

	count := len(symbols)

	for i := 0; i < count && len(symbols) != 0; i++ {

		switch symbols[i] {
		case "*":
			vectors = removeVector(vectors, i, '*', numbers)
			symbols = removeSymbols(symbols, i)
			_, numbers = numbers[0], numbers[1:]
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
