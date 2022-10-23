package rsa

import (
	"math/big"
	"strings"
)

var (
	rsa   *RSA
	split = "\xff\xfe\xff"
)

func Init() (err error) {
	rsa, err = New(0)
	return
}

func Encrypt(bs []byte, pub *PublicKey) (c []byte) {
	return rsa.Encrypt(bs, pub)
}

func Decrypt(be []byte, pri *PrivateKey) (bs []byte) {
	return rsa.Decrypt(be, pri)
}

type RSA struct {
	p uint64
	q uint64
	o uint64

	N uint64
	E uint64
	D uint64

	pub PublicKey
	pri PrivateKey
}

func New(N uint64) (*RSA, error) {
	p, q, err := generate2PrimeNumbers(N)
	if err != nil {
		return nil, err
	}

	n := p * q
	fi := euler(p, q)
	rsa := &RSA{p: p, q: q, o: fi, N: n}

	tps := sieve(rsa.o)
	tpsm := make(map[uint64]struct{}, len(tps))
	for _, t := range tps {
		tpsm[t] = struct{}{}
	}
	for t := range tpsm {
		if gcd(t, p-1) == 1 && gcd(t, q-1) == 1 {
			rsa.E = t
			break
		}
	}

	for i := rsa.o / rsa.E; i < rsa.o; i++ {
		if rsa.E*i%rsa.o == 1 {
			rsa.D = i
			break
		}
	}

	rsa.pub = PublicKey{N: rsa.N, E: rsa.E}
	rsa.pri = PrivateKey{N: rsa.N, D: rsa.D}

	return rsa, nil
}

func (r RSA) Encrypt(bs []byte, pub *PublicKey) (be []byte) {
	e, n := r.E, r.N
	if pub != nil && pub.Check() == nil {
		e = pub.E
		n = pub.N
	}

	bet := make([]string, 0)
	for _, b := range bs {
		m := new(big.Int).SetBytes([]byte{b})
		c := new(big.Int).Exp(m, big.NewInt(int64(e)), big.NewInt(int64(n)))
		bet = append(bet, string(c.Bytes()))
	}

	return []byte(strings.Join(bet, split))
}

func (r RSA) Decrypt(be []byte, pri *PrivateKey) (bs []byte) {
	d, n := r.D, r.N
	if pri != nil && pri.Check() == nil {
		d = pri.D
		n = pri.N
	}

	bs = make([]byte, 0)
	for _, b := range strings.Split(string(be), split) {
		c := new(big.Int).SetBytes([]byte(b))
		m := new(big.Int).Exp(c, big.NewInt(int64(d)), big.NewInt(int64(n)))
		bs = append(bs, m.Bytes()[0])
	}

	return bs
}

func (r RSA) PublicKey() string {
	return r.pub.String()
}

func (r RSA) PrivateKey() string {
	return r.pri.String()
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
