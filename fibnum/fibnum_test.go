package fibnum

import "testing"

type testFibCase struct {
	n    int64
	fibn int64
}

var testFibCases = []testFibCase{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{7, 13},
	{10, 55},
	{20, 6765},
	{30, 832040},
	{40, 102334155},
	{50, 12586269025},
	{60, 1548008755920},
	{70, 190392490709135},
	{80, 23416728348467685},
	{90, 2880067194370816120},
	{92, 7540113804746346429},
}

func TestFib1(t *testing.T) {
	for _, test := range testFibCases {
		if test.n < 50 { // very slow func
			f := Fib1(test.n)
			if f != test.fibn {
				t.Errorf("n = %v, expected: %v, got: %v", test.n, test.fibn, f)
			}
		}
	}
}

func TestFib2(t *testing.T) {
	for _, test := range testFibCases {
		f := Fib2(test.n)
		if f != test.fibn {
			t.Errorf("n = %v, expected: %v, got: %v", test.n, test.fibn, f)
		}
	}
}

func TestFib3(t *testing.T) {
	for _, test := range testFibCases {
		f := Fib3(test.n)
		if f != test.fibn {
			t.Errorf("n = %v, expected: %v, got: %v", test.n, test.fibn, f)
		}
	}
}
