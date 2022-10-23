package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/migregal/bmstu-iu7-ds/lab_04/pkg/rsa"

	"crypto/rand"
	rsa2 "crypto/rsa"
)

var (
	file, output, priKey, pubKey     string
	toDecrypt, toGenerate bool
)

func init() {
	flag.BoolVar(&toDecrypt, "d", false, "decrypt file")
	flag.BoolVar(&toGenerate, "g", false, "generate keys")
	flag.StringVar(&file, "f", "Makefile", "file to perform operation on")
	flag.StringVar(&output, "o", "rsaed", "file to store result")
	flag.StringVar(&priKey, "pri", "key_rsa", "private key")
	flag.StringVar(&pubKey, "pub", "key_rsa.pub", "public key")
}

func main() {
	flag.Parse()

	rsa2.GenerateKey(rand.Reader, 2048)

	if toGenerate && toDecrypt {
		log.Fatalln("Can't generate and decrypt file at the same time")
	}

	if toGenerate {
		key, err := rsa.New(0)
		if err != nil {
			log.Fatalf("failed to generate rsa: %s", err)
		}

		priGen := key.String()
		if err := os.WriteFile(priKey, []byte(priGen), 0644); err != nil {
			log.Fatalf("failed to write private key, error is: %s", err)
		}

		pubGen := key.PublicKey.String()
		if err := os.WriteFile(pubKey, []byte(pubGen), 0644); err != nil {
			log.Fatalf("failed to write public key, error is: %s", err)
		}
		return
	}

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Can't open file, error is: %s", err)
	}

	if toDecrypt {
		priFile, err := os.ReadFile(priKey)
		if err != nil {
			log.Fatalf("Can't open private key, error is: %s", err)
		}

		priParams := strings.Split(string(priFile), ",")
		priN, _ := strconv.Atoi(priParams[0])
		priD, _ := strconv.Atoi(priParams[1])
		privateKey := rsa.NewPrivateKey(uint64(priN), uint64(priD))

		decrypted := rsa.Decrypt(data, privateKey)

		if err := os.WriteFile(output, decrypted, 0644); err != nil {
			log.Fatalf("Can't write decrypted data, error is: %s", err)
		}
		return
	}

	pubFile, err := os.ReadFile(pubKey)
	if err != nil {
		log.Fatalf("Can't open public key, error is: %s", err)
	}

	pubParams := strings.Split(string(pubFile), ",")
	pubN, _ := strconv.Atoi(pubParams[0])
	pubE, _ := strconv.Atoi(pubParams[1])
	publicKey := rsa.NewPublicKey(uint64(pubN), uint64(pubE))

	encrypted := rsa.Encrypt(data, publicKey)

	if err := os.WriteFile(output, encrypted, 0644); err != nil {
		log.Fatalf("Can't write encrypted data, error is: %s", err)
	}
}
