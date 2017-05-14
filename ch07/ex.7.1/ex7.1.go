package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {

	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*wc += WordCounter(count)
	return count, nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {

	c := 0
	for _, b := range p {
		if b == '\n' {
			c++
		}
	}

	*lc += LineCounter(c)

	return 1, nil
}

func main() {

	var w WordCounter
	w.Write([]byte("Hello World! Whats up!"))
	fmt.Println(w)

	var lc LineCounter
	lc.Write([]byte(`hello
	abc
	def ad fas
	`))
	fmt.Println(lc)

}
