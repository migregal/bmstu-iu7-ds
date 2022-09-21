package enigma

type reflector struct {
	Values []rune
}

func newReflector(length int) *reflector {
	return &reflector{
		Values: fillReflector(length),
	}
}

func fillReflector(length int) []rune {
	values := make([]rune, length+1)
	for i := 0; i <= length; i++ {
		values[i] = rune(length - i)
	}
	return values
}
