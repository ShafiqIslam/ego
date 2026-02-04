package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}

func readerExercise() {
	fmt.Println("An exercise of reader to read infinite stream of character")

	reader.Validate(MyReader{})
}

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(p []byte) (int, error) {
	n, err := rr.r.Read(p)

	for i := 0; i < n; i++ {
		b := p[i]

		switch {
		case 'a' <= b && b <= 'z':
			p[i] = 'a' + (b-'a'+13)%26
		case 'A' <= b && b <= 'Z':
			p[i] = 'A' + (b-'A'+13)%26
		}
	}

	return n, err
}

func anotherReaderExercise() {
	fmt.Println("Another reader exercise with rot13 cipher")

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func helloReader() {
	fmt.Println("io.Reader interface represents the read end of a stream of data")

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func main() {
	helloReader()
	fmt.Println()
	readerExercise()
	fmt.Println()
	anotherReaderExercise()
}
