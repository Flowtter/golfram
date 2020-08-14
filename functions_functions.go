package main

import (
	"log"
	"strings"
)

func recognizeFunc(expression string) []string {
	// using x->x+2 format
	function := strings.Split(expression, "->")

	if len(function) != 2 {
		log.Fatal("Function recognize: bad input")
	}
	return function
}

func replaceFunc(expression string, element rune) []string {
	function := recognizeFunc(expression)
	count := len(function[1])

	r := []rune(function[0])
	t := r[0]

	r = []rune(function[1])

	for i := 0; i < count; i++ {
		if r[i] == t {
			r[i] = element
		}
	}
	return []string{function[0], string(r)}
}

func simplifyFunc(expression string, element rune) float64 {
	function := replaceFunc(expression, element)
	return basics(function[1])
}
