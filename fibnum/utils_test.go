package fibnum

import (
	"reflect"
	"testing"
)

func TestValidMatrixForMult(t *testing.T) {
	type testCase struct {
		a, b   [][]int64
		result bool
	}

	testCases := []testCase{
		{nil, nil, false},
		{[][]int64{}, [][]int64{}, false},
		{[][]int64{{1}, {1, 1}}, [][]int64{{1, 1}, {1, 1}}, false},
		{[][]int64{{1, 1}, {1, 1}}, [][]int64{{1, 1}, {1, 1, 1}}, false},
		{[][]int64{{1, 1}, {1, 1}}, [][]int64{{}, {1, 1}}, false},
		{[][]int64{{1, 1}, {}}, [][]int64{{}, {1, 1}}, false},
		{[][]int64{{1, 1}, {1, 1}}, [][]int64{{1, 1}, {1, 1}}, true},
		{[][]int64{{1, 1}, {2, 2}, {3, 3}}, [][]int64{{1, 1, 1}, {2, 2, 2}}, true},
		{[][]int64{{1, 1, 1}, {2, 2, 2}}, [][]int64{{1, 1}, {2, 2}}, false},
		{[][]int64{{1, 1, 1}, {2, 2, 2}}, [][]int64{{1}, {1}, {1}}, true},
	}

	for i, test := range testCases {
		valid := validMatrixForMult(test.a, test.b)
		if valid != test.result {
			t.Errorf("test = %v, expected: %v, got: %v", i, test.result, valid)
		}
	}
}

func TestMultMatrix(t *testing.T) {
	type testCase struct {
		a, b, c    [][]int64
		expetError bool
	}

	testCases := []testCase{
		{nil, nil, nil, true},
		{[][]int64{{1, 1}, {1, 1}, {1, 1}}, [][]int64{{2, 2}, {2, 2}, {2, 2}}, nil, true},
		{[][]int64{{1, 1}, {1, 1}}, [][]int64{{2, 2}, {2, 2}}, [][]int64{{4, 4}, {4, 4}}, false},
		{[][]int64{{1, 1}, {1, 1}, {1, 1}}, [][]int64{{2, 2, 2}, {2, 2, 2}}, [][]int64{{4, 4, 4}, {4, 4, 4}, {4, 4, 4}}, false},
		{[][]int64{{1, 1}, {1, 1}, {1, 1}}, [][]int64{{2}, {2}}, [][]int64{{4}, {4}, {4}}, false},
		{[][]int64{{1, 1}, {1, 1}}, [][]int64{{1, 0}, {0, 1}}, [][]int64{{1, 1}, {1, 1}}, false},
	}

	for i, test := range testCases {
		c, err := multMatrix(test.a, test.b)
		if err != nil && !test.expetError {
			t.Errorf("test = %v, expected no error, got: %v", i, err)
		} else if err == nil && test.expetError {
			t.Errorf("test = %v, expected error", i)
		}

		if !reflect.DeepEqual(c, test.c) {
			t.Errorf("test = %v, expected: %v, got: %v", i, test.c, c)
		}
	}
}
