package digsig

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func Verify(pubKey, file, sig string) (bool, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return false, err
	}

	key, err := os.ReadFile(pubKey)
	if err != nil {
		return false, err
	}

	signature, err := os.ReadFile(sig)
	if err != nil {
		return false, err
	}

	pem, _ := pem.Decode(key)
	pubKeyParsed, err := x509.ParsePKIXPublicKey(pem.Bytes)
	if err != nil {
		return false, err
	}

	return ed25519.Verify(pubKeyParsed.(ed25519.PublicKey), data, signature), nil
}
