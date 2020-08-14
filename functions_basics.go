package main

import (
	"strconv"
	"strings"
)

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

func slicerNumbersLetters(expression string) []string {
	match := strings.FieldsFunc(expression, split)
	var result = []string{}

	for _, i := range match {
		result = append(result, i)
	}
	return result
}

func split(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
}

/*
func slicerSymbols(expression string) []string {
	return strings.FieldsFunc(expression, splitNeg)
}

func splitNeg(r rune) bool {
	return r != '+' && r != '-' && r != '*' && r != '/'
}

*/

func slicerSymbols(expression string) []string {
	result := []string{}
	for _, element := range expression {
		if element == '+' || element == '-' || element == '*' || element == '/' {
			result = append(result, string(element))
		}
	}
	return result
}

func containsList(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsString(expression string, elements []rune) bool {
	for _, e := range elements {
		for _, a := range expression {
			if a == e {
				return true
			}
		}
	}
	return false
}

func removeNumber(slice []float64, s int, t rune) []float64 {
	if (s + 1) < len(slice) {
		if t == '*' {
			slice[s] *= slice[s+1]
		} else if t == '/' {
			slice[s] /= slice[s+1]
		} else if t == '+' {
			slice[s] += slice[s+1]
		} else if t == '-' {
			slice[s] -= slice[s+1]
		}

		return append(slice[:s+1], slice[s+2:]...)
	}
	return []float64{slice[0]}
}

func removeSymbols(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func checkParenthesis(expression string) bool {
	return containsString(expression, []rune{'(', ')'})
}

func handleExpressionParenthesis(expression string, position int) string {
	smallExpression := ""
	count := len(expression)
	positionLastParenthesis := count

	for i := position + 1; i < count; i++ {
		if expression[i] != ')' {
			smallExpression += string(expression[i])
		} else {
			positionLastParenthesis = i
			break
		}
	}

	result := basics(smallExpression)

	end := ""

	for i := 0; i < position; i++ {
		end += string(expression[i])
	}
	end += strconv.FormatFloat(result, 'f', -1, 64)

	for i := positionLastParenthesis + 1; i < count; i++ {
		end += string(expression[i])
	}
	return end
}

func handleParenthesis(expression string) string {

	open, openMax, lastOpenPosition, count := 0, 0, 0, len(expression)

	for i := 0; i < count; i++ {
		if expression[i] == '(' {
			open++
			if open > openMax {
				openMax = open
				lastOpenPosition = i
			}
		} else if expression[i] == ')' {
			open--
		}
	}

	return handleExpressionParenthesis(expression, lastOpenPosition)
}
