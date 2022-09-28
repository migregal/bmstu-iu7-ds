package enigma

import (
	"math/rand"
)

type reflector struct {
	values []byte

	from uint8
	to   uint8
}

func newReflector(seed int64, from, to uint8) *reflector {
	return &reflector{
		values: fillReflector(seed, from, to),
		from: from,
		to: to,
	}
}

func fillReflector(seed int64, from, to uint8) []byte {
	rand.Seed(seed)

	length := to - from

	stable := -1
	if length%2 != 0 {
		stable = rand.Int() % int(length)
	}

	values := make([]byte, length)
	for i := range values {
		values[i] = length
	}
	for i := uint8(0); i < length; i++ {
		if stable != -1 && i == uint8(stable) {
			values[i] = i
			continue
		}

		if values[i] != length {
			continue
		}

		x := uint8(rand.Uint32() % uint32(length))

		for (stable != -1 && x == uint8(stable)) || values[x] != length {
			x = uint8(rand.Uint32() % uint32(length))
		}
		values[i], values[x] = x, i
	}

	return values
}

func (r reflector) Reflect(b byte) byte {
	if b < r.from || b > r.to {
		return b
	}

	return r.values[b-r.from]+r.from
}
