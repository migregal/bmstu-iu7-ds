package enigma

import (
	"math/rand"
)

type reflector struct {
	values []byte
}

func newReflector(seed int64, length int) *reflector {
	return &reflector{
		values: fillReflector(seed, length),
	}
}

func fillReflector(seed int64, length int) []byte {
	rand.Seed(seed)

	stable := -1
	if length%2 != 0 {
		stable = rand.Int() % (length)
	}

	values := make([]byte, length)
	for i := range values {
		values[i] = byte(length)
	}
	for i := 0; i < length; i++ {
		if i == stable {
			values[i] = byte(i)
			continue
		}

		if values[i] != byte(length) {
			continue
		}

		x := rand.Intn(length)

		for x == stable || values[x] != byte(length) {
			x = rand.Intn(length)
		}
		values[i], values[x] = byte(x), byte(i)
	}

	return values
}

func (r reflector) Reflect(b byte) byte {
	return r.values[b]
}
