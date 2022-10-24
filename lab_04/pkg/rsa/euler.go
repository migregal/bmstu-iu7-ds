package rsa

import "math/big"

var bigOne = big.NewInt(1)

func euler(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(
		new(big.Int).Sub(a, bigOne),
		new(big.Int).Sub(b, bigOne))
}
