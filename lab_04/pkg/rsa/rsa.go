package rsa

import (
	"crypto/rand"
	"crypto/subtle"
	"fmt"
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

func Encrypt(msg []byte, key *PublicKey) ([]byte, error) {
	k := key.Size()

	if len(msg) > k-11 {
		return nil, fmt.Errorf("message is too long")
	}

	em := make([]byte, k)
	em[1] = 2
	ps, mm := em[2:len(em)-len(msg)-1], em[len(em)-len(msg):]
	err := nonZeroRandomBytes(ps, rand.Reader)
	if err != nil {
		return nil, err
	}
	em[len(em)-len(msg)-1] = 0
	copy(mm, msg)

	m := new(big.Int).SetBytes(em)
	c := new(big.Int).Exp(m, big.NewInt(int64(key.E)), key.N)

	return c.FillBytes(em), nil
}

func Decrypt(ciphertext []byte, key *PrivateKey) ([]byte, error) {
	k := key.Size()
	if k < 11 {
		return nil, fmt.Errorf("failed to decrypt: invalid key size")
	}

	c := new(big.Int).SetBytes(ciphertext)
	m := new(big.Int).Exp(c, key.D, key.N)

	em := m.FillBytes(make([]byte, k))

	var index int

	firstByteIsZero := subtle.ConstantTimeByteEq(em[0], 0)
	secondByteIsTwo := subtle.ConstantTimeByteEq(em[1], 2)

	// The remainder of the plaintext must be a string of non-zero random
	// octets, followed by a 0, followed by the message.
	//   lookingForIndex: 1 iff we are still looking for the zero.
	//   index: the offset of the first zero byte.
	lookingForIndex := 1

	for i := 2; i < len(em); i++ {
		equals0 := subtle.ConstantTimeByteEq(em[i], 0)
		index = subtle.ConstantTimeSelect(lookingForIndex&equals0, i, index)
		lookingForIndex = subtle.ConstantTimeSelect(equals0, 0, lookingForIndex)
	}

	// The PS padding must be at least 8 bytes long, and it starts two
	// bytes into em.
	validPS := subtle.ConstantTimeLessOrEq(2+8, index)

	valid := firstByteIsZero & secondByteIsTwo & (^lookingForIndex & 1) & validPS
	index = subtle.ConstantTimeSelect(valid, index+1, 0)

	if valid == 0 {
		panic("invalid hash")
	}
	return em[index:], nil
}
