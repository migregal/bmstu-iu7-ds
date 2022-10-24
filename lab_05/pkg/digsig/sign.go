package digsig

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func Sign(privKey, file, sig string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	key, err := os.ReadFile(privKey)
	if err != nil {
		return err
	}

	pem, _ := pem.Decode(key)
	privKeyParsed, err := x509.ParsePKCS8PrivateKey(pem.Bytes)
	if err != nil {
		return err
	}

	return os.WriteFile(sig, ed25519.Sign(privKeyParsed.(ed25519.PrivateKey), data), 0644)
}
