package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	status, err := bufio.NewReader(conn).ReadString(0)
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	fmt.Println(status)
}
