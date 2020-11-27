package arithmetic

import "testing"

func TestModExp(t *testing.T) {
	type testCase struct {
		x, y, N, p int64
	}

	testCases := []testCase{
		{3, 2, 5, 4},
		{1, 12, 7, 1},
		{12, 0, 13, 1},
		{3, 1300, 13, 3},
		{4, 13, 497, 445},
	}

	for i, test := range testCases {
		m := ModExp(test.x, test.y, test.N)
		if m != test.p {
			t.Errorf("test:%v, expected: %v, got: %v", i, test.p, m)
		}
	}
}

func TestEuclid(t *testing.T) {
	type testCase struct {
		a, b, d int64
	}

	testCases := []testCase{
		{12, 8, 4},
		{291, 71, 1},
		{35, 80, 5},
		{3571, 2357, 1},
	}

	for i, test := range testCases {
		m := Euclid(test.a, test.b)
		if m != test.d {
			t.Errorf("test:%v, expected: %v, got: %v", i, test.d, m)
		}
	}
}

func TestExtendedEuclid(t *testing.T) {
	type testCase struct {
		a, b, d, x, y int64
	}

	testCases := []testCase{
		{25, 11, 1, 4, -9},
	}

	for i, test := range testCases {
		x, y, d := ExtendedEuclid(test.a, test.b)
		if x != test.x || y != test.y || d != test.d {
			t.Errorf("test:%v, expected: %v, %v, %v, got: %v, %v, %v",
				i, test.x, test.y, test.d, x, y, d)
		}
	}
}

func TestPrimality(t *testing.T) {
	type testCase struct {
		N            int64
		Prim         bool
		IsCarmichael bool
	}

	testCases := []testCase{
		{71, true, false},
		{138, false, false},
		{561, false, true},
	}

	for i, test := range testCases {
		prim := Primality(test.N)
		if prim != test.Prim {
			if !test.IsCarmichael {
				t.Errorf("test:%v, expected: %v, got: %v, number: %v",
					i, test.Prim, prim, test.N)
			} else {
				// you are not lucky
			}
		}
	}
}

func TestPrimality2(t *testing.T) {
	type testCase struct {
		N            int64
		Prim         bool
		IsCarmichael bool
	}

	testCases := []testCase{
		{71, true, false},
		{138, false, false},
		{561, false, true},
	}

	for i, test := range testCases {
		prim := Primality2(test.N)
		if prim != test.Prim {
			if !test.IsCarmichael {
				t.Errorf("test:%v, expected: %v, got: %v, number: %v",
					i, test.Prim, prim, test.N)
			} else {
				t.Errorf("VERY RARE: test:%v, expected: %v, got: %v, number: %v",
					i, test.Prim, prim, test.N)
			}
		}
	}
}
