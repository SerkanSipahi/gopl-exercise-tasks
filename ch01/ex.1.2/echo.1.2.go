package main

import (
	"fmt"
	"os"
)

func main(){

	/**
	 * @Howto:
	 * 1.) go build echo.1.2.go
	 * 2.) ./echo.1.2 your args here
	 *
	 * @Notes:
	 * 1.) with "cmd + x" or "ctrl + x" for stopping the program
	 */

	for index, value := range os.Args[1:] {
		fmt.Println(index, value)
	}

}