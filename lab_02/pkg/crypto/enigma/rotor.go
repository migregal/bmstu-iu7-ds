package enigma

import (
	"math/rand"
)

type rotor struct {
	values    []rune
	rotations int
}

func newRotor(seed int64, length int) *rotor {
	return &rotor{
		values: fillRotor(seed, length),
	}
}

func fillRotor(seed int64, length int) []rune {
	values := make([]rune, length+1)
	for i := 0; i <= length; i++ {
		values[i] = rune(i)
	}

	rand.Seed(seed)
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	return values
}

func (r *rotor) rotate() bool {
	r.values = append([]rune{r.values[len(r.values)-1]}, r.values[:len(r.values)-1]...)
	r.rotations = (r.rotations + 1) % len(r.values)

	return r.rotations == 0
}

func (r *rotor) getStraight(b rune) rune {
	return r.values[b]
}

func (r *rotor) getReverse(b rune) rune {
	for i, v := range r.values {
		if v == b {
			return rune(i)
		}
	}

	return 0
}
