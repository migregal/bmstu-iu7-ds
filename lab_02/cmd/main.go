package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"math"
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
	alphabetSize := math.MaxUint16
	enigmaMachine := enigma.NewEnigma(seed, alphabetSize)

	encoded := enigmaMachine.Cipher([]rune(string(data)))
	if err = os.WriteFile(fout, []byte(string(encoded)), 0644); err != nil {
		log.Fatalln(err)
	}

	enigmaMachine = enigma.NewEnigma(seed, alphabetSize)
	decoded := enigmaMachine.Decipher(encoded)
	if err = os.WriteFile(check, []byte(string(decoded)), 0644); err != nil {
		log.Fatalln(err)
	}

	isEqual := bytes.Equal(data, []byte(string(decoded)))
	fmt.Printf("Input and decoded files are equal: %t\n", isEqual)
}
