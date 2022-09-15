package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"time"

	"github.com/briandowns/spinner"
	"github.com/logrusorgru/aurora"

	"github.com/migregal/bmstu-iu7-ds/pkg/keygen"
)

func main() {
	fmt.Println(aurora.BgBlue("Performing installation of utility..."))
	s := spinner.New(spinner.CharSets[32], 100*time.Millisecond)
	s.Color("bgYellow", "bold", "fgBlack")
	s.Start()
	defer func() {
		time.Sleep(5 * time.Second)
		s.Stop()
	}()

	key, err := keygen.GetKey()
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile("./app.out")
	if err != nil {
		log.Fatal(err)
	}

	reg, err := regexp.Compile(keygen.KeyRegexp);
	if err != nil {
		log.Fatal(err)
	}


	data = reg.ReplaceAll(data, []byte(key));

	ioutil.WriteFile("./app.out", data, 0744)
}
