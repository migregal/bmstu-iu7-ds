package rsa

import (
	"math/big"
	"strings"
)

var (
	split = "\xff\xfe\xff"
)

func New(N uint64) (*PrivateKey, error) {
	p, q, err := generate2PrimeNumbers(N)
	if err != nil {
		return nil, err
	}

	var key PrivateKey
	key.N = p * q

	fi := euler(p, q)

	tps := sieve(fi)
	tpsm := make(map[uint64]struct{}, len(tps))
	for _, t := range tps {
		tpsm[t] = struct{}{}
	}
	for t := range tpsm {
		if gcd(t, p-1) == 1 && gcd(t, q-1) == 1 {
			key.E = t
			break
		}
	}

	for i := fi / key.E; i < fi; i++ {
		if key.E*i%fi == 1 {
			key.D = i
			break
		}
	}

	return &key, nil
}

func Encrypt(bs []byte, key *PublicKey) (be []byte) {
	if key == nil || key.Check() != nil {
		return nil
	}

	bet := make([]string, 0)
	for _, b := range bs {
		m := new(big.Int).SetBytes([]byte{b})
		c := new(big.Int).Exp(m, big.NewInt(int64(key.E)), big.NewInt(int64(key.N)))
		bet = append(bet, string(c.Bytes()))
	}

	return []byte(strings.Join(bet, split))
}

func Decrypt(be []byte, key *PrivateKey) (bs []byte) {
	if key == nil || key.Check() != nil {
		return nil
	}

	bs = make([]byte, 0)
	for _, b := range strings.Split(string(be), split) {
		c := new(big.Int).SetBytes([]byte(b))
		m := new(big.Int).Exp(c, big.NewInt(int64(key.D)), big.NewInt(int64(key.N)))
		bs = append(bs, m.Bytes()[0])
	}

	return bs
}

func gcd(m, n uint64) uint64 {
	if m < n {
		m, n = n, m
	}

	if n == 0 {
		return m
	}

	return gcd(n, m%n)
}
