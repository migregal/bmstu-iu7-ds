package main

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/migregal/bmstu-iu7-ds/lab-01/pkg/keygen"
)

const licenseKey = "00000000-0000-0000-0000-000000000000"

func main() {
	isLicensed, err := keygen.CheckKey(licenseKey)
	if err != nil {
		fmt.Printf("%s %s\n", aurora.BgRed("Error while checking license key:"), err)
		os.Exit(1)
	}

	if !isLicensed {
		fmt.Println(aurora.BgRed("Program is not registered for this PC. Aborting"))
		os.Exit(1)
	}

	fmt.Printf("Hello, {{ %s }}\n", licenseKey)
}
