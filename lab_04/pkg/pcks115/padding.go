package pcks115

import (
	"crypto/rand"
	"crypto/subtle"
)

var MinKeySize = 11

func Padding(src []byte, k int) ([]byte, error) {
	em := make([]byte, k)
	em[1] = 2
	ps, mm := em[2:len(em)-len(src)-1], em[len(em)-len(src):]
	err := nonZeroRandomBytes(ps, rand.Reader)
	if err != nil {
		return nil, err
	}
	em[len(em)-len(src)-1] = 0
	copy(mm, src)

	return em, nil
}

func Unpadding(src []byte) []byte {
	var index int

	firstByteIsZero := subtle.ConstantTimeByteEq(src[0], 0)
	secondByteIsTwo := subtle.ConstantTimeByteEq(src[1], 2)

	// The remainder of the plaintext must be a string of non-zero random
	// octets, followed by a 0, followed by the message.
	//   lookingForIndex: 1 iff we are still looking for the zero.
	//   index: the offset of the first zero byte.
	lookingForIndex := 1

	for i := 2; i < len(src); i++ {
		equals0 := subtle.ConstantTimeByteEq(src[i], 0)
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

	return src[index:]
}
