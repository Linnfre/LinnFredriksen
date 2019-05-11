package fileutils

import (
	"bufio"
	"io"
	"log"
	"os"
)

func FileToByteslice(filename string) []byte {

	// Open file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	return byteSlice

}

func Lineshift(filename string) string {
	var lineending string
	// Open file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(line) > 1 {
			lineending = line[len(line)-2:]
			break
		}
	}

	if lineending[0] < 31 {
		return lineending
	} else {
		return lineending[1:]
	}
}
