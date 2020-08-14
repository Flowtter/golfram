package main

import "testing"

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
				t.Errorf("TestVectorsAdd of (%s + %s) was incorrect, got: %s, want: %s.", table.vector1, table.vector2, total, table.result)
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
				t.Errorf("TestVectorsSubstract of (%s + %s) was incorrect, got: %s, want: %s.", table.vector1, table.vector2, total, table.result)
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

func TestVectorsSplit(t *testing.T) {
	tables := []struct {
		expression string
		result     []string
	}{

		{"(2,6)", []string{"2", "6"}},
		{"(-2,6,6)", []string{"-2", "6", "6"}},
		{"(x,-y,6)", []string{"x", "-y", "6"}},
		{"(x+1,2-y,6)", []string{"x+1", "2-y", "6"}},
	}

	for _, table := range tables {
		total := splitVector(table.expression)
		count := len(total)
		for i := 0; i < count; i++ {
			if total[i] != table.result[i] {
				t.Errorf("TestVectorsSplit of (%s) was incorrect, got: %s, want: %s.", table.expression, total, table.result)
				break
			}
		}

	}
}
