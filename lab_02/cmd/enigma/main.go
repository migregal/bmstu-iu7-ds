package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/migregal/bmstu-iu7-ds/lab-02/pkg/crypto/enigma"
)

var (
	fin, fout, check string
)

func init() {
	flag.StringVar(&fin, "input", "go.mod", "file to encode")
	flag.StringVar(&fout, "output", "go.mod.encoded", "encoding result file")
	flag.StringVar(&check, "check", "go.mod.decoded", "decoding encoded file result for clearance")
}

func main() {
	flag.Parse()

	data, err := os.ReadFile(fin)
	if err != nil {
		log.Fatalln(err)
	}

	seed := time.Now().UnixNano()
	enigmaMachine := enigma.DefaultEnigma(seed)

	encoded := enigmaMachine.Cipher(data)
	if err = os.WriteFile(fout, encoded, 0644); err != nil {
		log.Fatalln(err)
	}

	enigmaMachine = enigma.DefaultEnigma(seed)
	decoded := enigmaMachine.Decipher(encoded)
	if err = os.WriteFile(check, []byte(string(decoded)), 0644); err != nil {
		log.Fatalln(err)
	}

	isEqual := bytes.Equal(data, []byte(string(decoded)))
	fmt.Printf("Input and decoded files are equal: %t\n", isEqual)
}
