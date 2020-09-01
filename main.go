package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func response(c *gin.Context, statusCode int, err error, payload interface{}) {
	var status string
	if statusCode >= 200 && statusCode <= 299 {
		status = "success"
	} else {
		status = "error"
	}
	data := gin.H{
		"code":    statusCode,
		"status":  status,
		"payload": payload,
	}
	if err != nil {
		data["error"] = err.Error()
	}

	c.JSON(statusCode, data)
}

func main() {
	r := gin.Default()
	r.GET("/request/:name", func(c *gin.Context) {
		name := c.Param("name")
		name = strings.ReplaceAll(name, "§", "/")
		requested := request(name)
		response(c, http.StatusOK, nil, requested)
	})

	r.Use(static.Serve("/", static.LocalFile("html", true)))
	r.NoRoute(func(c *gin.Context) {
		c.File("./html/index.html")
	})

	r.Run()
}

func request(expression string) string {

	text := strings.Replace(expression, "\n", "", -1)

	text = strings.ReplaceAll(text, ` `, ``)

	parse := strings.Split(text, `:`)

	if len(parse) == 0 {
		return "Please parse an expression"
	} else if len(parse) == 1 {
		/*
			if parse[0] == "help" || parse[0] == "h" {
				return `Welcome to Golfram \n
				To parse an expression, use semicolons. Space will be erased by Golfram, use them if it can help you. (2 + 3) == (2+3)
				basic, b, and an expression to solve it. Example : \" basic: 2 + 3 / 4 * 5 \" \n
				vector, v, to work on vectors. Example : \" vector : (2,y)*z+(x,6)*2 \" \n
				degree, d, to get the degree of an expression. Example : \" degree : x->x^3+x^2+x+1 \" \n
				function, f, to get the result of the function with a parameter. Example : \" function : x->x+2 : 2 \" \n`
			} else {
		*/
		return "Please parse an expression"
		//}
	} else if len(parse) == 2 {
		if parse[0] == "basic" || parse[0] == "b" {
			return strconv.FormatFloat(basics(parse[1]), 'f', -1, 64)
		} else if parse[0] == "vector" || parse[0] == "v" {
			return returnVector(basicsVectors(parse[1]))
		} else if parse[0] == "degree" || parse[0] == "d" {
			return strconv.Itoa(getDegree(parse[1]))
		} else {
			return "Please parse a correct argument"
		}

	} else if len(parse) == 3 {
		if parse[0] == "function" || parse[0] == "f" {
			return strconv.FormatFloat(simplifyFunc(parse[1], parse[2]), 'f', -1, 64)
		} else {
			return "Please parse correct arguments"
		}

	}
	return "Please parse an expression (2 arguments max)"
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
