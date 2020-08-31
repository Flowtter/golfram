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
	parse := strings.Fields(text)

	return strconv.FormatFloat(basics(parse[1]), 'f', -1, 64)
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
