package rsa

import "math/big"

var bigOne = big.NewInt(1)

func euler(primes []*big.Int) *big.Int {
	totient := new(big.Int).Set(bigOne)
	pminus1 := new(big.Int)
	for _, prime := range primes {
		pminus1.Sub(prime, bigOne)
		totient.Mul(totient, pminus1)
	}

	return totient
}
