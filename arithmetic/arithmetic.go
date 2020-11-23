package arithmetic

// Multiply - Musa al-Khwarizmi's algorithm
// T(n) = O(n^2), where n = Min(log2(a),log2(b))
func Multiply(a, b int64) int64 {
	if a == 0 || b == 0 {
		return 0
	}

	// try to reduce the number of cycle steps
	// e.g. a = 1073741824 b = 2 -> so instead of log2(a)=30 steps
	// we have log2(b) = 1 step
	x := min(a, b)
	y := max(a, b)
	var mult int64
	for ; x > 0; x = x >> 1 { // in fact, we n times shift
		if x%2 == 1 {
			mult += y // O(n)
		}
		y = y << 1 // O(1) (in registers of the processor this operation needs O(1))
	}
	// O(n) + ... + O(n) (n times) = n * O(n) = O(n^2)
	return mult
}

// Divide - returns q, r
// a = q*b + r
func Divide(a, b int64) (q, r int64) {
	if b == 0 {
		panic(b)
	}
	if a == 0 {
		return 0, 0
	}
	q, r = Divide(a/2, b)
	q = q << 1 // q *= 2
	r = r << 1 // r *= 2
	if a%2 == 1 {
		r++
	}
	if r >= b {
		r -= b
		q++
	}
	return q, r
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
