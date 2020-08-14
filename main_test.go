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

func TestVectorsBasics(t *testing.T) {
	tables := []struct {
		expression string
		result     []string
	}{

		{"(2,6)", []string{"2", "6"}},
		{"(2,6)*2", []string{"4", "12"}},
		{"(2,6)+(1,3)", []string{"3", "9"}},
		{"(2,6)-(2,6)", []string{"0", "0"}},
		{"(2,6)-(2,3)", []string{"0", "3"}},
		{"(2,3)-(2,6)", []string{"0", "-3"}},
		{"(2,6)*2+(2,6)", []string{"6", "18"}},
		{"(2,6)*2+(1,3)*3", []string{"7", "21"}},

		{"(2,6)*x", []string{"2*x", "6*x"}},
		{"(2,y)-(x,6)", []string{"2-x", "y-6"}},
		{"(2,y)+(x,6)", []string{"2+x", "y+6"}},
		{"(2,y)*z+(x,6)*2", []string{"2*z+x*2", "y*z+12"}},
	}

	for _, table := range tables {
		total := basicsVectors(table.expression)
		count := len(total)
		for i := 0; i < count; i++ {
			if total[i] != table.result[i] {
				t.Errorf("TestVectorsSplit of (%s) was incorrect, got: %s, want: %s.", table.expression, total, table.result)
				break
			}
		}

	}
}
