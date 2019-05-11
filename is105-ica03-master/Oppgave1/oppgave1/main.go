package main

import (
	"fmt"

	fileutils "github.com/LinnFredriksen/is105-ica03-master/Oppgave1/fileutils"
)

func main() {
	file1 := fileutils.FileToByteslice("text1.txt")
	contents := fileutils.FileToByteslice("text2.txt")
	fmt.Println("Text1", file1)
	fmt.Println()
	fmt.Println("Text2", contents)
}
