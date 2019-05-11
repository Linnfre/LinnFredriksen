package main

import (
	"fmt"
	"os"

	unicode "github.com/LinnFredriksen/is105-ica02-master/unicode/uni"
)

func main() {
	//Skriv jp eller is etter 'go run unicode_main.go'
	if len(os.Args) > 1 {
		language := os.Args[1]
		expression := "nord og sør"
		trans := unicode.Translate(expression, language)
		fmt.Printf("%s\n", trans)
	} else {
		fmt.Println("Ingen språk spesifisert!")
	}
}
