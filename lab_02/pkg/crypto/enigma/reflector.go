package enigma

import (
	"math/rand"
)

type reflector struct {
	values []rune
}

func newReflector(seed int64, length int) *reflector {
	return &reflector{
		values: fillReflector(seed, length),
	}
}

func fillReflector(seed int64, length int) []rune {
	rand.Seed(seed)

	stable := -1
	if length%2 != 0 {
		stable = rand.Int() % (length)
	}

	values := make([]rune, length)
	for i := range values {
		values[i] = rune(length)
	}
	for i := 0; i < length; i++ {
		if i == stable {
			values[i] = rune(i)
			continue
		}

		if values[i] != rune(length) {
			continue
		}

		x := rand.Intn(length)

		for x == stable || values[x] != rune(length) {
			x = rand.Intn(length)
		}
		values[i], values[x] = rune(x), rune(i)
	}

	return values
}

func (r reflector) Reflect(b rune) rune {
	return r.values[b]
}
