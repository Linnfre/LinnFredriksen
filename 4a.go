package main

import (

	"fmt"
	"os"
	"./unicode"
)

func main() {

	language := os.Args[1]
	expression := "nord og sør"
	trans := unicode.Oversetter(expression, language)
	fmt.Print("%s", trans)

}

