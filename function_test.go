package main

import "testing"

func TestRecognize(t *testing.T) {
	tables := []struct {
		expression string
		result     []string
	}{

		{"x->x+2", []string{"x", "x+2"}},
		{"y->y-2", []string{"y", "y-2"}},
		{"z->z*2", []string{"z", "z*2"}},
		{"z->z*2+4", []string{"z", "z*2+4"}},
	}

	for _, table := range tables {
		total := recognizeFunc(table.expression)
		count := len(total)
		for i := 0; i < count; i++ {
			if total[i] != table.result[i] {
				t.Errorf("TestRecognize of (%s) was incorrect, got: %s, want: %s.", table.expression, total, table.result)
				break
			}
		}
	}
}

func TestReplace(t *testing.T) {
	tables := []struct {
		expression string
		element    rune
		result     []string
	}{

		{"x->x+2", '2', []string{"x", "2+2"}},
		{"y->y-2", '4', []string{"y", "4-2"}},
		{"z->z*2", '3', []string{"z", "3*2"}},
		{"z->z*2+4", '4', []string{"z", "4*2+4"}},
	}

	for _, table := range tables {
		total := replaceFunc(table.expression, table.element)
		count := len(total)
		for i := 0; i < count; i++ {
			if total[i] != table.result[i] {
				t.Errorf("TestReplace of (%s) was incorrect, got: %s, want: %s.", table.expression, total, table.result)
				break
			}
		}
	}
}
func TestSimplify(t *testing.T) {
	tables := []struct {
		expression string
		element    rune
		result     float64
	}{

		{"x->x+2", '2', 4},
		{"y->y-2", '4', 2},
		{"z->z*2", '3', 6},
		{"z->z*2+4", '4', 12},
		{"z->2*z^2+z+3", '4', 39},
	}

	for _, table := range tables {
		total := simplifyFunc(table.expression, table.element)
		if total != table.result {
			t.Errorf("TestReplace of (%s) was incorrect, got: %f, want: %f.", table.expression, total, table.result)
		}
	}
}
