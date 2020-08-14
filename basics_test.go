package main

import "testing"

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
