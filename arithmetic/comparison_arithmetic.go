package arithmetic

import "math/rand"

// ModExp return x^y mod N
func ModExp(x, y, N int64) int64 {
	if y == 0 {
		return 1
	}
	z := ModExp(x, y/2, N)
	t := (z * z) % N
	if y%2 == 0 {
		return t
	}
	return (x * t) % N
}

// Euclid return d = gcd(a,b)
func Euclid(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return Euclid(b, a%b)
}

// ExtendedEuclid return d = gcd(a,b) and ax + by = d
func ExtendedEuclid(a, b int64) (x, y, d int64) {
	if b == 0 {
		return 1, 0, a
	}
	xn, yn, d := ExtendedEuclid(b, a%b)
	return yn, xn - (a/b)*yn, d
}

// Primality //
func Primality(N int64) bool {
	a := rand.Int63n(N)
	if ModExp(a, N-1, N) == 1 {
		return true
	}
	return false
}

// Primality2 //
func Primality2(N int64) bool {
	k := 20
	random := make([]int64, 0, k)
	for i := 0; i < k; i++ {
		random = append(random, rand.Int63n(N))
	}

	isPrim := true
	for i := range random {
		isPrim = isPrim && (ModExp(random[i], N-1, N) == 1)
	}

	return isPrim
}
