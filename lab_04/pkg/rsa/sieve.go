package rsa

import "math"

func sieve(n uint64) (ps []uint64) {
	ps = make([]uint64, 0)
	if n < 2 {
		return ps
	}

	N := make([]bool, n+1)
	for i, l := uint64(2), uint64(math.Sqrt(float64(n))); i <= l; i++ {
		if !N[i] {
			for j := uint64(2); i*j <= n; j++ {
				N[i*j] = true
			}
		}
	}

	for i, l := uint64(2), n+1; i < l; i++ {
		if !N[i] {
			ps = append(ps, i)
		}
	}

	return ps
}
