package rsa

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"

	"github.com/migregal/bmstu-iu7-ds/lab_04/pkg/pcks115"
)

func New(rnd io.Reader, bits int) (*PrivateKey, error) {
	var key PrivateKey
	key.E = 65537

	for {
		// N = p * q
		p, err := rand.Prime(rnd, bits/2)
		if err != nil {
			return nil, err
		}
		q, err := rand.Prime(rnd, (bits - p.BitLen()))
		if err != nil {
			return nil, err
		}

		if p.Cmp(q) == 0 {
			continue
		}

		key.N = new(big.Int).Mul(p, q)
		if key.N.BitLen() != bits {
			continue
		}

		// fi = (p - 1) * (q - 1)
		fi := euler(p, q)

		key.D = new(big.Int)
		e := big.NewInt(int64(key.E))
		ok := key.D.ModInverse(e, fi)

		if ok != nil {
			break
		}
	}

	return &key, nil
}

func Encrypt(key *PublicKey, in io.Reader, out io.Writer) error {
	k := key.Size()
	if k < pcks115.MinKeySize {
		return fmt.Errorf("invalid key")
	}

	for {
		buf := make([]byte, k-pcks115.MinKeySize)
		n, err := in.Read(buf)
		if err == io.EOF {
			break
		}

		em, err := pcks115.Padding(buf[:n], k)
		if err != nil {
			return err
		}

		buf = make([]byte, k)

		m := new(big.Int).SetBytes(em)
		c := new(big.Int).Exp(m, big.NewInt(int64(key.E)), key.N)
		c.FillBytes(buf)

		out.Write(buf)
	}

	return nil
}

func Decrypt(key *PrivateKey, in io.Reader, out io.Writer) error {
	k := key.Size()
	if k < pcks115.MinKeySize {
		return fmt.Errorf("failed to decrypt: invalid key size")
	}

	buf := make([]byte, k)
	for {
		_, err := in.Read(buf)
		if err == io.EOF {
			break
		}

		c := new(big.Int).SetBytes(buf)
		m := new(big.Int).Exp(c, key.D, key.N)

		em := m.FillBytes(make([]byte, k))

		out.Write(pcks115.Unpadding(em))
	}

	return nil
}
