package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	buf := make([]byte, 4)
	for {
		n, err := r.Read(buf)
		fmt.Println(n, err, buf[:n], string(buf[:n]))
		if err == io.EOF {
			break
		}

	}

	// use io.ReadFull to read exactly len(buf) bytes into buf
	r2 := strings.NewReader("abcde")
	buf2 := make([]byte, 4)

	if _, err2 := io.ReadFull(r2, buf2); err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(buf2, string(buf2))
}
