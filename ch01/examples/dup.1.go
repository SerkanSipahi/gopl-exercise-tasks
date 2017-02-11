package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){

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