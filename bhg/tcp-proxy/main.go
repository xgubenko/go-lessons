package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FooReader struct{}
type FooWriter struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

func (FooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to copy data")
	}

	// input := make([]byte, 4096)

	// s, err := reader.Read(input)

	// if err != nil {
	// 	log.Fatalln("Unable to read data")
	// }
	// fmt.Printf("Read %d bytes from stdin\n", s)

	// s, err = writer.Write(input)

	// if err != nil {
	// 	log.Fatalln("Unable to write data")
	// }

	// fmt.Printf("Wrote %d bytes to stdout\n", s)
}
