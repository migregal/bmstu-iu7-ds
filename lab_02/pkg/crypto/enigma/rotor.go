package enigma

import (
	"math/rand"
)

type rotor struct {
	values    []rune
	reverse   []rune
	rotations int
}

func newRotor(seed int64, length int) *rotor {
	value, reverse := fillRotor(seed, length)

	return &rotor{
		values: value,
		reverse: reverse,
	}
}

func fillRotor(seed int64, length int) ([]rune, []rune) {
	values := make([]rune, length)
	for i := 0; i < length; i++ {
		values[i] = rune(i)
	}

	rand.Seed(seed)
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	reverse := make([]rune, length)
	for i := 0; i < length; i++ {
		reverse[values[i]] = rune(i)
	}

	return values, reverse
}

func (r *rotor) rotate() bool {
	r.rotations = (r.rotations + 1) % len(r.values)
	return r.rotations == 0
}

func (r *rotor) getStraight(b rune) rune {
	idx := (r.rotations + int(b)) % len(r.values)
	return rune((int(r.values[idx]) - r.rotations + len(r.values))%len(r.values))
}

func (r *rotor) getReverse(b rune) rune {
	idx := (r.rotations + int(b)) % len(r.reverse)
	return rune((int(r.reverse[idx]) - r.rotations + len(r.reverse))%len(r.reverse))
}
