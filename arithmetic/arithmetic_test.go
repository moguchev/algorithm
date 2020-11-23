package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	type testCase struct {
		a, b int64
	}

	testCases := []testCase{
		{0, 0},
		{0, 1},
		{1, 0},
		{11, 13},
		{20, 10},
		{128, 2},
		{12812812812812, 3},
	}

	for i, test := range testCases {
		m := Multiply(test.a, test.b)
		if m != test.a*test.b {
			t.Errorf("test:%v, expected: %v, got: %v", i, test.a*test.b, m)
		}
	}
}

func TestDivide(t *testing.T) {
	type testCase struct {
		a, b int64
	}

	testCases := []testCase{
		{64, 2},
		{123, 14},
		{0, 12314},
		{12812812812812, 12812812812812},
	}

	for i, test := range testCases {
		q, r := Divide(test.a, test.b)
		if q != test.a/test.b || r != test.a%test.b {
			t.Errorf("test:%v, expected: (%v,%v), got: (%v,%v)", i, test.a/test.b, test.a%test.b, q, r)
		}
	}
}

func TestDividePanic(t *testing.T) {
	assert.Panics(t, func() { Divide(1, 0) }, "expected panic")
}
