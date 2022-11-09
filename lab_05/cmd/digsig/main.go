package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/migregal/bmstu-iu7/lab_05/pkg/digsig"
)

var (
	generate        bool
	privKey, pubKey string

	sign      bool
	file, sig string
)

func init() {
	flag.BoolVar(&generate, "g", false, "generate keys for dig sig")
	flag.StringVar(&privKey, "priv", "privkey.pem", "private key filename")
	flag.StringVar(&pubKey, "pub", "pubkey.pem", "public key filename")

	flag.BoolVar(&sign, "s", false, "sign file keys with private key")
	flag.StringVar(&file, "f", "test.txt", "file to sign")
	flag.StringVar(&sig, "sig", "signature.sig", "file to save signature")
}

func main() {
	flag.Parse()

	if generate && sign {
		log.Fatal("invalid args: can't generate and sign at the same time")
	}

	if generate {
		generateKeys()
		return
	}

	if sign {
		signData()
		return
	}

	verify()
}

func generateKeys() {
	privateKeyPEM, err := os.Create(privKey)
	if err != nil {
		log.Fatal(err)
	}
	defer privateKeyPEM.Close()

	publicKeyPEM, err := os.Create(pubKey)
	if err != nil {
		log.Fatal(err)
	}
	defer publicKeyPEM.Close()

	if err := digsig.GenerateKeys(privateKeyPEM, publicKeyPEM); err != nil {
		log.Fatalf("Failed to generate keys, error is: %s", err)
	}
}

func signData() {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	key, err := os.ReadFile(privKey)
	if err != nil {
		log.Fatal(err)
	}

	signature, err := digsig.Sign(data, key)
	if err != nil {
		log.Fatalf("Failed to sign file, error is: %s", err)
	}

	if err := os.WriteFile(sig, signature, 0644); err != nil {
		log.Fatal(err)
	}
}

func verify() {
	isValid, err := digsig.Verify(pubKey, file, sig)
	if err != nil {
		log.Fatalf("Failed to sign file, error is: %s", err)
	}
	if !isValid {
		fmt.Println("Signature is corrupted")
	} else {
		fmt.Println("Signature is correct")
	}
}
