package enigma

import (
	"math"
)

const defaultRotorsCount = 3

type Enigma struct {
	rotors    []*rotor
	reflector *reflector
}

func DefaultEnigma(seed int64) Enigma {
	return NewEnigma(seed, 0, math.MaxUint8, defaultRotorsCount)
}

func NewEnigma(seed int64, from, to uint8, rotorsN int) Enigma {
	rotors := make([]*rotor, rotorsN)
	for i := 0; i < rotorsN; i++ {
		rotors[i] = newRotor(seed, from, to)
	}

	return Enigma{
		rotors:    rotors,
		reflector: newReflector(seed, from, to),
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
