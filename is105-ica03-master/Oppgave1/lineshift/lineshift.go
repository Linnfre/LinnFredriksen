package main

import (
	"fmt"
	"os"

	fileutils "github.com/LinnFredriksen/is105-ica03-master/Oppgave1/fileutils"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		filename := args[1]
		lineending := fileutils.Lineshift(filename)
		if lineending == "\n" {
			fmt.Printf("Line ending of %s is LF (%+q)\n", filename, lineending)
		} else if lineending == "\r\n" {
			fmt.Printf("Line ending of %s is CRLF (%+q)\n", filename, lineending)
		}
	} else {
		fmt.Println("Usage: lineshift [filename]")
	}
}
