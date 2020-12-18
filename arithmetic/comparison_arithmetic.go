package arithmetic

import (
	"fmt"
	"math"
	"math/rand"
)

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
		rnd := rand.Int63n(N)
		if rnd > 1 {
			random = append(random, rnd)
		} else {
			i--
		}
	}

	for i := range random {
		if ModExp(random[i], N-1, N) != 1 {

		}
	}

	return true
}

// BreakeRSA //
func BreakeRSA(N, e, C int64) (p, q, M int64) {
	fmt.Println(int64(math.Sqrt(float64(N))))
	var i int64
	for i = 2; i <= int64(math.Sqrt(float64(N))); i++ {
		if N%i == 0 {
			if Primality2(i) {
				p = i
				q = N / p
				fmt.Printf("p = %d, q = %d\n", p, q)
			}
		}
	}

	fin := (p - 1) * (q - 1)
	fmt.Printf("fi(n) = %d\n", fin)

	d, _, _ := ExtendedEuclid(e, fin)
	if d < 0 {
		d += fin
	}
	fmt.Printf("d = %d\n", d)

	M = ModExp(C, d, N)
	fmt.Printf("M = %d\n", M)

	fmt.Printf("Binary M: %b\n", M)

	return p, q, M
}
