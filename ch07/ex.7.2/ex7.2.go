package main

import (
	"fmt"
	"io"
	"os"
)

// Write a function CountingWriter with the the signature(*1) below that, given an
// io.Writer, returns a new Writer(*2) that wraps the original (*3), and a pointer(*4) to an
// int64 variable that at any moment contains the number of bytes written to
// to the new Writer.

type ByteCounter struct {
	w     io.Writer // (*3)
	count int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {

	n, err := c.w.Write(p)
	c.count += int64(n)

	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) { // (*1)

	c := ByteCounter{w, 0} // (*2)

	return &c, &c.count // (*4)
}

func main() {

	writer, count := CountingWriter(os.Stdout)
	fmt.Fprint(writer, "Hello World\n")
	fmt.Println(*count)
}
