package rsa

import (
	"crypto/rand"
	"io"
	"math/big"
)

func New(rnd io.Reader, bits int) (*PrivateKey, error) {
	var key PrivateKey
	key.E = 65537

	var err error

	for {
		// N = p * q
		key.N = new(big.Int).SetInt64(1)
		for i := 0; i < 2; i++ {
			var tmp *big.Int
			tmp, err = rand.Prime(rnd, bits/(2-i))
			if err != nil {
				return nil, err
			}

			key.primes = append(key.primes, tmp)
			key.N = key.N.Mul(key.N, tmp)
		}

		// fi = (p - 1) * (q - 1)
		fi := euler(key.primes)

		key.D = new(big.Int)
		e := big.NewInt(int64(key.E))
		ok := key.D.ModInverse(e, fi)

		if ok != nil {
			break
		}
	}

	return &key, nil
}

func Encrypt(bet []byte, key *PublicKey) (be []byte) {
	m := new(big.Int).SetBytes(bet)
	c := new(big.Int).Exp(m, big.NewInt(int64(key.E)), key.N)

	return c.Bytes()
}

func Decrypt(ciphertext []byte, key *PrivateKey) (bs []byte) {
	c := new(big.Int).SetBytes(ciphertext)
	m := new(big.Int).Exp(c, key.D, key.N)

	return m.Bytes()
}
