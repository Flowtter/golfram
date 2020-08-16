package main

import (
	"log"
	"strconv"
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

func getDegree(expression string) int {
	function := recognizeFunc(expression)

	xL := []rune(function[0])
	x := xL[0]

	if function[1] == "0" {
		return -1
	}
	if !containsString(function[1], xL) {
		return 0
	}
	highest := 1

	count := len(function[1])

	decomposedFunction := []rune(function[1])
	recording := false
	actual := ""

	for i := 0; i < count; i++ {
		if decomposedFunction[i] == x && i+2 < count && decomposedFunction[i+1] == '^' {
			recording = true
			i += 2
		}
		if recording && decomposedFunction[i] >= '0' && decomposedFunction[i] <= '9' {
			actual += string(decomposedFunction[i])
		} else {
			a, err := strconv.Atoi(actual)
			if err == nil && a > highest {
				highest = a
			}
			actual = ""
			recording = false
		}
	}
	if actual != "" {
		a, err := strconv.Atoi(actual)
		if err == nil && a > highest {
			highest = a
		}
	}

	return highest
}
