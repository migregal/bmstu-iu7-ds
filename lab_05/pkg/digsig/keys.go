package digsig

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func GenerateKeys(privKey, pubKey string) error {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}

	err = generatePrivateKey(privKey, privateKey)
	if err != nil {
		return err
	}

	err = generatePublicKey(pubKey, publicKey)
	if err != nil {
		return err
	}

	return nil
}

func generatePrivateKey(key string, privateKey ed25519.PrivateKey) error {
	privateKeyPEM, err := os.Create(key)
	if err != nil {
		return err
	}
	defer privateKeyPEM.Close()

	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return err
	}

	if err := pem.Encode(privateKeyPEM, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privBytes,
	}); err != nil {
		return err
	}

	return nil
}

func generatePublicKey(key string, publicKey ed25519.PublicKey) error {
	publicKeyPEM, err := os.Create(key)
	if err != nil {
		return err
	}
	defer publicKeyPEM.Close()

	pubBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	if err := pem.Encode(publicKeyPEM, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}); err != nil {
		return err
	}

	return nil
}
