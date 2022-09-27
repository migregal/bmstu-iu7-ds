package enigma

import (
	"math"
)

const defaultMaxRotorsCount = 8

type Enigma struct {
	rotors    []*rotor
	reflector *reflector
}

func DefaultEnigma(seed int64) Enigma {
	return NewEnigma(seed, math.MaxUint8, defaultMaxRotorsCount)
}

func NewEnigma(seed int64, alphabetLen int, rotorsN int) Enigma {
	rotors := make([]*rotor, rotorsN)
	for i := 0; i < rotorsN; i++ {
		rotors[i] = newRotor(seed, alphabetLen)
	}

	return Enigma{
		rotors:    rotors,
		reflector: newReflector(seed, alphabetLen),
	}
}

func (e *Enigma) Cipher(data []byte) []byte {
	ciphered := make([]byte, len(data))
	for i, b := range data {
		ciphered[i] = e.encodebyte(b)
	}

	return ciphered
}

func (e *Enigma) encodebyte(data byte) byte {
	for i := range e.rotors {
		data = e.rotors[i].getStraight(data)
	}

	data = e.reflector.Reflect(data)

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

func (e *Enigma) Decipher(data []byte) []byte {
	return e.Cipher(data)
}
