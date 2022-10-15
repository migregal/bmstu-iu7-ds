package main

import (
	"flag"
	"log"
	"os"

	"github.com/migregal/bmstu-iu7-ds/lab-03/pkg/crypto/des"
	stddes "github.com/migregal/bmstu-iu7-ds/lab-03/pkg/crypto/des/std"
)

var (
	file, output, key string
	decrypt, std      bool
)

func init() {
	flag.BoolVar(&decrypt, "d", false, "decrypt file")
	flag.StringVar(&file, "f", "Makefile", "file to perform operation on")
	flag.StringVar(&output, "o", "des.enc", "file to store result")
	flag.StringVar(&key, "k", "password", "encryption/decryption key")
	flag.BoolVar(&std, "std", false, "use stdlib realization")
}

func main() {
	flag.Parse()

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Can't open file, error is: %s", err)
	}

	var out []byte

	key = des.CompleteKey(key)
	if decrypt {
		if std {
			out, err = stddes.Decipher([]byte(key), []byte(key), data)
		} else {
			out, err = des.Decipher(key, data)
		}
		if err != nil {
			log.Fatalf("Can't decrypt data, error is: %s", err)
		}

		if err := os.WriteFile(output, out, 0644); err != nil {
			log.Fatalf("Can't write decrypted data, error is: %s", err)
		}
	} else {
		if std {
			out, err = stddes.Cipher([]byte(key), []byte(key), data)
		} else {
			out, err = des.Cipher(key, data)
		}
		if err != nil {
			log.Fatalf("Can't encrypt data, error is: %s", err)
		}

		if err := os.WriteFile(output, out, 0644); err != nil {
			log.Fatalf("Can't write encrypted data, error is: %s", err)
		}
	}
}
