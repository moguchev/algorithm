package fibnum

// Fib1 : T(n) = O(2^n), M(n) = O(1) <- stack memory don`t considered
func Fib1(n int64) int64 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Fib1(n-1) + Fib1(n-2)
	}
}

// Fib2 : T(n) = O(n), M(n) = O(n)
func Fib2(n int64) int64 {
	if n == 0 {
		return 0
	}
	series := make([]int64, n+1, n+1)
	series[0] = 0
	series[1] = 1
	for i := 2; i < len(series); i++ {
		series[i] = series[i-1] + series[i-2]
	}
	return series[n]
}

// Fib3 :  T(n) = O(logn), M(n) = O(1)
// (Fn+1 Fn  ) = (1 1)^n
// (Fn   Fn-1)   (1 0)
func Fib3(n int64) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	count := [][]int64{{1, 0}, {0, 1}}

	m := [][]int64{{1, 1}, {1, 0}}

	// fast pow matrix
	for p := n; p > 0; {
		if p%2 == 0 {
			p /= 2
			// m *= m
			m, _ = multMatrix(m, m)
		} else {
			p--
			//count *= m
			count, _ = multMatrix(count, m)
		}
	}

	return count[0][1]
}
