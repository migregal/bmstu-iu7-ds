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
	if k < 11 {
		return fmt.Errorf("invalid key")
	}

	for {
		buf := make([]byte, k-11)
		n, err := in.Read(buf)
		if err == io.EOF {
			break
		}

		// PCKS1_v1_15
		em := make([]byte, k)
		em[1] = 2
		ps, mm := em[2:len(em)-n-1], em[len(em)-n:]
		err = nonZeroRandomBytes(ps, rand.Reader)
		if err != nil {
			return err
		}
		em[len(em)-n-1] = 0
		copy(mm, buf)

		// em := buf[:n]
		buf = make([]byte, k)

		m := new(big.Int).SetBytes(em)
		c := new(big.Int).Exp(m, big.NewInt(int64(key.E)), key.N)
		c.FillBytes(buf)

		// em = append(bytes.Repeat([]byte{0}, k-len(em)), em...)
		out.Write(buf)
	}

	// m := new(big.Int).SetBytes(em)
	// c := new(big.Int).Exp(m, big.NewInt(int64(key.E)), key.N)

	// return c.FillBytes(em), nil

	return nil
}

func Decrypt(key *PrivateKey, in io.Reader, out io.Writer) error {
	k := key.Size()
	if k < 11 {
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

		out.Write(em[index:])
	}

	return nil
}
