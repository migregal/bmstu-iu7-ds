package enigma

import (
	"math/rand"
)

type rotor struct {
	values    []byte
	reverse   []byte
	rotations int

	from uint8
	to   uint8
}

func newRotor(seed int64, from, to uint8) *rotor {
	value, reverse := fillRotor(seed, from, to)

	return &rotor{
		values:  value,
		reverse: reverse,
		from:    from,
		to:      to,
	}
}

func fillRotor(seed int64, from, to uint8) ([]byte, []byte) {
	values := make([]byte, to-from)
	for i := from; i < to; i++ {
		values[i-from] = i - from
	}

	rand.Seed(seed)
	rand.Shuffle(int(to-from), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	reverse := make([]byte, to-from)
	for i := from; i < to; i++ {
		reverse[values[i-from]] = i - from
	}

	return values, reverse
}

func (r *rotor) rotate() bool {
	r.rotations = (r.rotations + 1) % len(r.values)
	return r.rotations == 0
}

func (r *rotor) getStraight(b byte) byte {
	return extractFromSlice(b, r.values, r.from,r.to,r.rotations)
}

func (r *rotor) getReverse(b byte) byte {
	return extractFromSlice(b, r.reverse, r.from,r.to,r.rotations)
}

func extractFromSlice(b byte, data []byte, from, to uint8, rotations int) byte {
	if b < from || b > to {
		return b
	}

	idx := (rotations + int(b-from)) % len(data)
	return byte((int(data[idx])-rotations+len(data))%len(data)) + from
}
