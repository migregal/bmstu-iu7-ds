package main

import (
	"flag"
	"log"
	"os"

	"github.com/migregal/bmstu-iu7-ds/lab-03/pkg/crypto/des"
)

var (
	file, output, key string
	decrypt           bool
)

func init() {
	flag.BoolVar(&decrypt, "d", false, "decrypt file")
	flag.StringVar(&file, "f", "Makefile", "file to perform operation on")
	flag.StringVar(&output, "o", "desed", "file to store result")
	flag.StringVar(&key, "k", "password", "encryption/decryption key")
}

func main() {
	flag.Parse()

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Can't open file, error is: %s", err)
	}

	key = des.CompleteKey(key)
	if decrypt {
		decrypted := des.Decrypt(data, des.GenerateKeys(key))
		if err := os.WriteFile(output, decrypted, 0644); err != nil {
			log.Fatalf("Can't write decrypted data, error is: %s", err)
		}
	} else {
		encrypted := des.Encrypt(data, des.GenerateKeys(key))
		if err := os.WriteFile(output, encrypted, 0644); err != nil {
			log.Fatalf("Can't write encrypted data, error is: %s", err)
		}
	}
}
