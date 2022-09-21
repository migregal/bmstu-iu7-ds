package crypto

type Cipherer interface {
	Encipher()
	Decipher()
}
