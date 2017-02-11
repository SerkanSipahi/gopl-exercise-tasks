package main

/**
 * @Howto:
 * 1.) go build echo.1.3.go
 * 2.) ./echo.1.3 your args here
 *
 * @Notes:
 * 1.) with "cmd + x" or "ctrl + x" for stopping the program
 */

import (
	"strings"
	"os"
	"time"
	"fmt"
	"bufio"
)

func Echo1(args []string) string {

	result := args[0] + " " + args[0]
	return result

}

func Echo2(args []string) string {

	result := strings.Join(args[0:], " ")
	return result

}

func Echo3(args []string) string {

	result := ""
	for index, value := range args[0:] {
		result += fmt.Sprintf("%d %s\n", index, value)
	}
	return result
}

func askForArguments(s string) []string {

	reader := bufio.NewReader(os.Stdin)
	args := []string{}

	for {

		fmt.Println(s)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		if response == "\n" {
			continue
		}

		response = strings.TrimRight(response, "\n")
		args = strings.Split(response, " ")
		if len(args) > 0 {
			break
		}

	}

	return args
}

func main(){

	args := os.Args
	if len(args[0:]) == 1 {
		args = askForArguments("Please provide at least one argument")
	}

	/**
	 * Book Section 1.6 with time package
	 */
	count := 5000000

	// Echo1
	start := time.Now()
	for i :=0; i < count; i++ {
		Echo1(args)
	}
	fmt.Printf("%.2fs\n", time.Since(start).Seconds())

	// Echo2
	start = time.Now()
	for i :=0; i < count; i++ {
		Echo2(args)
	}
	fmt.Printf("%.2fs\n", time.Since(start).Seconds())

	// Echo3
	start = time.Now()
	for i :=0; i < count; i++ {
		Echo3(args)
	}
	fmt.Printf("%.2fs\n", time.Since(start).Seconds())
}