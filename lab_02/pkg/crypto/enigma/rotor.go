package enigma

import (
	"math/rand"
)

type rotor struct {
	values    []byte
	reverse   []byte
	rotations int
}

func newRotor(seed int64, length int) *rotor {
	value, reverse := fillRotor(seed, length)

	return &rotor{
		values: value,
		reverse: reverse,
	}
}

func fillRotor(seed int64, length int) ([]byte, []byte) {
	values := make([]byte, length)
	for i := 0; i < length; i++ {
		values[i] = byte(i)
	}

	rand.Seed(seed)
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	reverse := make([]byte, length)
	for i := 0; i < length; i++ {
		reverse[values[i]] = byte(i)
	}

	return values, reverse
}

func (r *rotor) rotate() bool {
	r.rotations = (r.rotations + 1) % len(r.values)
	return r.rotations == 0
}

func (r *rotor) getStraight(b byte) byte {
	idx := (r.rotations + int(b)) % len(r.values)
	return byte((int(r.values[idx]) - r.rotations + len(r.values))%len(r.values))
}

func (r *rotor) getReverse(b byte) byte {
	idx := (r.rotations + int(b)) % len(r.reverse)
	return byte((int(r.reverse[idx]) - r.rotations + len(r.reverse))%len(r.reverse))
}
