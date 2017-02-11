package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main(){

	args := make([]string, 2)

	// when passed over cli
	if len(os.Args[1:]) > 0 {
		args = os.Args[1:]
	// when no args passed (e.g. running with an IDE). Useful when debugging
	} else {
		dir, _ := os.Getwd()
		args[0] = dir + "/" + "ch01/examples/file.1.txt"
		args[1] = dir + "/" + "ch01/examples/file.2.txt"
	}

	counts := map[string]int{}
	for _, filename := range args {

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}