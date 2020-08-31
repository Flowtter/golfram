package main

import (
	"fmt"
	"strconv"
	"strings"
)

func displayVector(vectors []string) {
	count := len(vectors)
	if count == 0 {
		return
	}
	count--
	fmt.Print("(")
	for i := 0; i < count; i++ {
		fmt.Print(vectors[i] + ", ")
	}
	fmt.Print(vectors[count])
	fmt.Println(")")
}

func returnVector(vectors []string) string {
	count := len(vectors)
	result := ""
	if count == 0 {
		return ""
	}
	count--
	result += "("
	for i := 0; i < count; i++ {
		result += vectors[i] + ", "
	}
	result += vectors[count]
	result += ")"
	return result
}

func splitVector(expression string) []string {
	expression = strings.Replace(expression, "(", "", -1)
	expression = strings.Replace(expression, ")", "", -1)
	s := strings.Split(expression, ",")
	return s
}

func addVectors(vector1 []string, vector2 []string) []string {
	count := len(vector1)
	if count != len(vector2) || count == 0 {
		fmt.Println("Add Vectors: vectors are not compatible")
		return []string{}
	}
	for i := 0; i < count; i++ {
		v1, err1 := strconv.ParseFloat(vector1[i], 64)
		v2, err2 := strconv.ParseFloat(vector2[i], 64)
		if err1 == nil && err2 == nil {
			vector1[i] = strconv.FormatFloat(v1+v2, 'f', -1, 64)
		} else {
			vector1[i] += "+" + vector2[i]
		}
	}
	return vector1
}

func substractVectors(vector1 []string, vector2 []string) []string {
	count := len(vector1)
	if count != len(vector2) || count == 0 {
		fmt.Println("Substract Vectors: vectors are not compatible")
		return []string{}
	}
	for i := 0; i < count; i++ {
		v1, err1 := strconv.ParseFloat(vector1[i], 64)
		v2, err2 := strconv.ParseFloat(vector2[i], 64)
		if err1 == nil && err2 == nil {
			vector1[i] = strconv.FormatFloat(v1-v2, 'f', -1, 64)
		} else {
			vector1[i] += "-" + vector2[i]
		}
	}
	return vector1
}

func homothetie(vector []string, lambda string) []string {
	count := len(vector)
	l, err2 := strconv.ParseFloat(lambda, 64)

	for i := 0; i < count; i++ {
		v1, err1 := strconv.ParseFloat(vector[i], 64)
		if err1 == nil && err2 == nil {
			vector[i] = strconv.FormatFloat(v1*l, 'f', -1, 64)
		} else {
			vector[i] += "*" + lambda
		}
	}
	return vector
}

func handleExpressionParenthesisVectors(expression string, position int) ([]string, string) {
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

	result := strings.Split(smallExpression, ",")

	end := ""

	for i := 0; i < position; i++ {
		end += string(expression[i])
	}
	for i := positionLastParenthesis + 1; i < count; i++ {
		end += string(expression[i])
	}
	return result, end
}

func handleParenthesisVectors(expression string) ([]string, string) {

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
	return handleExpressionParenthesisVectors(expression, lastOpenPosition)
}

func removeVector(vectors [][]string, s int, t rune, numbers []string) [][]string {
	if (s+1) < len(vectors) && t != '*' {
		if t == '+' {
			vectors[s] = addVectors(vectors[s], vectors[s+1])
		} else if t == '-' {
			vectors[s] = substractVectors(vectors[s], vectors[s+1])
		}

		return append(vectors[:s+1], vectors[s+2:]...)
	}
	vectors[s] = homothetie(vectors[s], numbers[0])

	return vectors
}
