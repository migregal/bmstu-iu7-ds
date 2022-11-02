package main

import (
	"flag"
	"log"
	"os"

	"github.com/migregal/bmstu-iu7-ds/lab_06/pkg/huffman"
)

var (
	file, outComp, outDecomp string
)

func init() {
	flag.StringVar(&file, "f", "Makefile", "file to compress")
	flag.StringVar(&outComp, "c", "compressed", "file to store compressed result")
	flag.StringVar(&outDecomp, "d", "decompressed", "file to store decompressed result")
}

func main() {
	flag.Parse()

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Can't open file, error is: %s", err)
	}

	tree := huffman.NewHuffmanTree(data)
	var compressed []byte
	if err := tree.Compress(&compressed); err != nil {
		log.Fatalf("Failed to compress, error is: %s", err)
	}
	if err := os.WriteFile(outComp, compressed, 0644); err != nil {
		log.Fatalf("Can't write compressed data, error is: %s", err)
	}
	deceompressed, err := tree.Decompress(compressed)
	if err != nil {
		log.Fatalf("Failed to decompress, error is: %s", err)
	}
	if err := os.WriteFile(outDecomp, deceompressed, 0644); err != nil {
		log.Fatalf("Can't write decompressed data, error is: %s", err)
	}
}
