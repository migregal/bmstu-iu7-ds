package rsa

import "fmt"

type PrivateKey struct {
	N uint64
	D uint64
}

func NewPrivateKey(n, d uint64) *PrivateKey {
	return &PrivateKey{n, d}
}

func (p PrivateKey) String() string {
	return fmt.Sprintf("%d,%d", p.N, p.D)
}

func (p PrivateKey) Check() error {
	if p.N < 6 || p.D < 2 {
		return fmt.Errorf("private key error, too small params")
	}
	return nil
}

type PublicKey struct {
	N uint64
	E uint64
}

func NewPublicKey(n, e uint64) *PublicKey {
	return &PublicKey{n, e}
}

func (p PublicKey) String() string {
	return fmt.Sprintf("%d,%d", p.N, p.E)
}

func (p PublicKey) Check() error {
	if p.N < 6 || p.E < 2 {
		return fmt.Errorf("public key error, too small params")
	}
	return nil
}
