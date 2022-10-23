package rsa

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type PrivateKey struct {
	PublicKey
	D *big.Int

	primes []*big.Int
}

func NewPrivateKey(data []byte) (*PrivateKey, error) {
	key := string(data)

	params := strings.Split(key, ",")
	if len(params) != 2 {
		return nil, fmt.Errorf("invalid params")
	}

	n, ok := new(big.Int).SetString(params[0], 0)
	if !ok {
		return nil, fmt.Errorf("invalid key structure")
	}

	d, ok := new(big.Int).SetString(params[1], 0)
	if !ok {
		return nil, fmt.Errorf("invalid key structure")
	}

	return &PrivateKey{PublicKey{N: n}, d, nil}, nil
}

func (p PrivateKey) String() string {
	return fmt.Sprintf("%s,%s", p.N.String(), p.D.String())
}

type PublicKey struct {
	N *big.Int
	E int
}

func NewPublicKey(data []byte) (*PublicKey, error) {
	key := string(data)
	params := strings.Split(key, ",")
	if len(params) != 2 {
		return nil, fmt.Errorf("invalid params")
	}

	n, ok := new(big.Int).SetString(params[0], 0)
	if !ok {
		return nil, fmt.Errorf("invalid key structure")
	}

	e, err := strconv.Atoi(params[1])
	if err != nil {
		return nil, fmt.Errorf("invalid key structure")
	}

	return &PublicKey{n, e}, nil
}

func (pub *PublicKey) Size() int {
	return (pub.N.BitLen() + 7) / 8
}

func (p PublicKey) String() string {
	return fmt.Sprintf("%s,%d", p.N.String(), p.E)
}
