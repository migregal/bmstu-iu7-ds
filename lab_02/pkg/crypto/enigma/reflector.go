package enigma

type reflector struct {
	values []rune
}

func newReflector(length int) *reflector {
	return &reflector{
		values: fillReflector(length),
	}
}

func fillReflector(length int) []rune {
	values := make([]rune, length+1)
	for i := 0; i <= length; i++ {
		values[i] = rune(length - i)
	}
	return values
}

func (r reflector) Reflect(b rune) rune {
	return r.values[b]
}
