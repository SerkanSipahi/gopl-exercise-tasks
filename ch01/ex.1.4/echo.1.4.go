package main

/**
 * @Howto:
 * 1.) go build echo.1.4.go
 * 2.) ./echo.1.4
 */

import (
	"os"
	"fmt"
	"bufio"
)

func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text() + " [ "+ f.Name() +" ]"]++
	}
}

func main(){

	counts := map[string]int{}
	files  := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {

			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}