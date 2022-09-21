package enigma

import (
	"math"
	"math/rand"
)

const maxRotorsCount = 8

type Enigma struct {
	rotors    []*rotor
	reflector *reflector
}

func DefaultEnigma(seed int64) Enigma {
	return NewEnigma(seed, math.MaxUint8)
}

func NewEnigma(seed int64, alphabetLen int) Enigma {
	rand.Seed(seed)

	size := rand.Int() % maxRotorsCount + 1
	rotors := make([]*rotor, size)
	for i := 0; i < size; i++ {
		rotors[i] = newRotor(seed, alphabetLen)
	}

	return Enigma{
		rotors:    rotors,
		reflector: newReflector(alphabetLen),
	}
}

func (e *Enigma) Cipher(data []rune) []rune {
	ciphered := make([]rune, len(data))
	for i, b := range data {
		ciphered[i] = e.encodeRune(b)
	}

	return ciphered
}

func (e *Enigma) encodeRune(data rune) rune {
	for i := range e.rotors {
		data = e.rotors[i].getStraight(data)
	}

	data = e.reflector.Values[data]

	for i := len(e.rotors) - 1; i >= 0; i-- {
		data = e.rotors[i].getReverse(data)
	}

	for _, rotor := range e.rotors {
		if !rotor.rotate() {
			break
		}
	}

	return data
}

func (e *Enigma) Decipher(data []rune) []rune {
	return e.Cipher(data)
}
