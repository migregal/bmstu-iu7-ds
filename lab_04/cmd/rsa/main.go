package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/migregal/bmstu-iu7-ds/lab_04/pkg/rsa"

	"crypto/rand"
)

var (
	file, output, priKey, pubKey string
	toDecrypt, toGenerate        bool
	keySize                      uint
)

func init() {
	flag.BoolVar(&toDecrypt, "d", false, "decrypt file")
	flag.BoolVar(&toGenerate, "g", false, "generate keys")
	flag.UintVar(&keySize, "s", 128, "key size in bytes")
	flag.StringVar(&file, "f", "Makefile", "file to perform operation on")
	flag.StringVar(&output, "o", "rsaed", "file to store result")
	flag.StringVar(&priKey, "pri", "key_rsa", "private key")
	flag.StringVar(&pubKey, "pub", "key_rsa.pub", "public key")
}

func main() {
	flag.Parse()

	if toGenerate && toDecrypt {
		log.Fatalln("can't generate and decrypt file at the same time")
	}

	if toGenerate {
		key, err := rsa.New(rand.Reader, int(keySize*8))
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

	fin, err := os.Open(file)
	if err != nil {
		log.Fatalf("can't open file, error is: %s", err)
	}
	defer fin.Close()

	in := bufio.NewReader(fin)

	fout, err := os.Create(output)
	if err != nil {
		log.Fatalf("can't open file, error is: %s", err)
	}
	defer fout.Close()

	out := bufio.NewWriter(fout)
	defer out.Flush()

	if toDecrypt {
		priFile, err := os.ReadFile(priKey)
		if err != nil {
			log.Fatalf("can't open private key, error is: %s", err)
		}

		privateKey, err := rsa.NewPrivateKey(priFile)
		if err != nil {
			log.Fatalf("can't parse private key, error is: %s", err)
		}

		err = rsa.Decrypt(privateKey, in, out)
		if err != nil {
			log.Fatalf("can't decrypt data, error is: %s", err)
		}

		return
	}

	pubFile, err := os.ReadFile(pubKey)
	if err != nil {
		log.Fatalf("can't open public key, error is: %s", err)
	}

	publicKey, err := rsa.NewPublicKey(pubFile)
	if err != nil {
		log.Fatalf("can't parse public key, error is: %s", err)
	}

	err = rsa.Encrypt(publicKey, in, out)
	if err != nil {
		log.Fatalf("can't encrypt data, error is: %s", err)
	}
}
