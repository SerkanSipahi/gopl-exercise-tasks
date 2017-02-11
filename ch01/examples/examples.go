package main

import (
	"bufio"
	"os"
	"fmt"
)

/**
 * @Howto:
 * 1.) go build examples.go
 * 2.) ./examples file.1.txt file.2.txt
 * or just call
 * 3.) ./examples
 *
 * @Notes:
 * 1.) with "cmd + x" or "ctrl + x" for stopping the program
 */

func Dup1(){

	/**
	 * map[string]int{} vs make(map[string]int)
	 * Is the same: look at here: https://blog.golang.org/go-maps-in-action
	 */

	counts := map[string]int{}
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func Dup2(){

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

func countLines(f *os.File, counts map[string]int){

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

}

func main(){

	// Dup1()
	Dup2()

}