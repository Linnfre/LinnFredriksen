package main

import (
	"fmt"
	"os"

	unicode "github.com/LinnFredriksen/is105-ica02-master/unicode/uni"
)

func main() {
	//Skriv jp eller is etter 'go run unicode_main.go'
	language := os.Args[1]
	fmt.Printf("%s\n", os.Args)
	expression := "nord og sør"
	trans := unicode.Translate(expression, language)
	fmt.Printf("%s", trans)
}
