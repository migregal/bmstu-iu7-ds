package digsig

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
)

func Sign(data, key []byte) ([]byte, error) {
	pem, _ := pem.Decode(key)
	privKeyParsed, err := x509.ParsePKCS8PrivateKey(pem.Bytes)
	if err != nil {
		return nil, err
	}

	return ed25519.Sign(privKeyParsed.(ed25519.PrivateKey), data), nil
}
