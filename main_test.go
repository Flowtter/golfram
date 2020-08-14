package main

import "testing"

func TestBasics(t *testing.T) {
	tables := []struct {
		expression string
		result     float64
	}{

		{"1+1", 2},
		{"1+2", 3},
		{"4+2+7", 13},

		{"1-1", 0},
		{"4-2", 2},
		{"4-2-7", -5},

		{"5+4-2", 7},
		{"5-10+2", -3},

		{"1*2", 2},
		{"4*2", 8},
		{"4*3*2", 24},

		{"4/2", 2},
		{"2/2", 1},
		{"1/2", 0.5},
		{"4/2/2/2", 0.5},

		{"4*2/8", 1},
		{"12/3*2", 8},

		{"2+2*3+2", 10},

		{"2*4+3*2-4/2", 12},
	}

	for _, table := range tables {
		total := basics(table.expression)
		if total != table.result {
			t.Errorf("TestBasics of (%s) was incorrect, got: %f, want: %f.", table.expression, total, table.result)
		}
	}
}

func TestParenthesis(t *testing.T) {
	tables := []struct {
		expression string
		result     float64
	}{

		{"(1+1)", 2},
		{"(1)+2", 3},
		{"4+(2+7)", 13},

		{"(1-1)", 0},
		{"(4)-2", 2},
		{"(4-2)-7", -5},

		{"(5+4)-2", 7},
		{"5-(10+2)", -7},

		{"(1*2)", 2},
		{"4*(2)", 8},
		{"4*(3*2)", 24},

		{"(4/2)", 2},
		{"(2)/2", 1},
		{"1/(2)", 0.5},
		{"4/2/(2/2)", 2},

		{"4*(2/8)", 1},
		{"12/(3*2)", 2},

		{"2+2*(3+2)", 12},
		{"2*(4+3)*2-4/2", 26},
	}

	for _, table := range tables {
		total := basics(table.expression)
		if total != table.result {
			t.Errorf("TestParenthesis of (%s) was incorrect, got: %f, want: %f.", table.expression, total, table.result)
		}
	}
}

func TestVectorsAdd(t *testing.T) {
	tables := []struct {
		vector1 []string
		vector2 []string
		result  []string
	}{

		{[]string{"1", "3"}, []string{"1", "3"}, []string{"2", "6"}},
		{[]string{"-1", "3"}, []string{"1", "3"}, []string{"0", "6"}},
		{[]string{"4", "3"}, []string{"2", "1"}, []string{"6", "4"}},
		{[]string{"1", "x"}, []string{"1", "3"}, []string{"2", "x+3"}},
		{[]string{"x", "y"}, []string{"z", "e"}, []string{"x+z", "y+e"}},
	}

	for _, table := range tables {
		total := addVectors(table.vector1, table.vector2)
		count := len(total)
		for i := 0; i < count; i++ {
			if total[i] != table.result[i] {
				t.Errorf("TestVectorsHomothetie of (%s + %s) was incorrect, got: %s, want: %s.", table.vector1, table.vector2, total, table.result)
				break
			}
		}

	}
}

func TestVectorsSubstract(t *testing.T) {
	tables := []struct {
		vector1 []string
		vector2 []string
		result  []string
	}{

		{[]string{"1", "3"}, []string{"1", "3"}, []string{"0", "0"}},
		{[]string{"-1", "3"}, []string{"1", "3"}, []string{"-2", "0"}},
		{[]string{"4", "3"}, []string{"2", "1"}, []string{"2", "2"}},
		{[]string{"1", "x"}, []string{"1", "3"}, []string{"0", "x-3"}},
		{[]string{"x", "y"}, []string{"z", "e"}, []string{"x-z", "y-e"}},
	}

	for _, table := range tables {
		total := substractVectors(table.vector1, table.vector2)
		count := len(total)
		for i := 0; i < count; i++ {
			if total[i] != table.result[i] {
				t.Errorf("TestVectorsHomothetie of (%s + %s) was incorrect, got: %s, want: %s.", table.vector1, table.vector2, total, table.result)
				break
			}
		}

	}
}

func TestVectorsHomothetie(t *testing.T) {
	tables := []struct {
		expression []string
		scalar     string
		result     []string
	}{

		{[]string{"1", "3"}, "2", []string{"2", "6"}},
		{[]string{"-1", "3"}, "2", []string{"-2", "6"}},
		{[]string{"1", "3"}, "x", []string{"1*x", "3*x"}},
	}

	for _, table := range tables {
		total := homothetie(table.expression, table.scalar)
		count := len(total)
		for i := 0; i < count; i++ {
			if total[i] != table.result[i] {
				t.Errorf("TestVectorsHomothetie of (%s) was incorrect, got: %s, want: %s.", table.expression, total, table.result)
				break
			}
		}

	}
}
