package fibnum

import "fmt"

func validMatrixForMult(a, b [][]int64) bool {
	if len(a) == 0 || len(b) == 0 { // empty matrixs
		return false
	}

	// check that a is matrix indeed
	columsNum := 0
	for i := range a {
		if i == 0 {
			columsNum = len(a[i])
		} else if len(a[i]) != columsNum || len(a[i]) == 0 {
			return false
		}
	}
	// check that b is matrix indeed
	columsNum = 0
	for i := range b {
		if i == 0 {
			columsNum = len(b[i])
		} else if len(b[i]) != columsNum || len(b[i]) == 0 {
			return false
		}
	}

	// number of columns A and number of rows B must be the same
	if len(a[0]) != len(b) {
		return false
	}

	return true
}

// C(m,n) = A(m, k) * B(k, n)
// Cij = ∑s=1,k(Ais*Bsj) i=1,m j = 1,n
// M(a,b) - matrix with a rows and b columns
func multMatrix(a, b [][]int64) ([][]int64, error) {
	if !validMatrixForMult(a, b) {
		return nil, fmt.Errorf("bad inputs")
	}

	m := len(a) // num rows of res
	k := len(b)
	n := len(b[0]) // num colums of res
	// create matrix
	result := make([][]int64, m, m)
	for i := range result {
		result[i] = make([]int64, n, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for s := 0; s < k; s++ {
				result[i][j] += (a[i][s] * b[s][j])
			}
		}
	}

	return result, nil
}
