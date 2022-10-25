package digsig

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"io"
)

func GenerateKeys(privKey, pubKey io.Writer) error {
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

func generatePrivateKey(out io.Writer, privateKey ed25519.PrivateKey) error {
	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return err
	}

	if err := pem.Encode(out, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privBytes,
	}); err != nil {
		return err
	}

	return nil
}

func generatePublicKey(out io.Writer, publicKey ed25519.PublicKey) error {
	pubBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	if err := pem.Encode(out, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}); err != nil {
		return err
	}

	return nil
}
