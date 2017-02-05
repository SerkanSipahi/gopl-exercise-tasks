package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main(){

	// Variant 1 ( can consume n arguments )
	fmt.Println(strings.Join(os.Args[0:], " "))

	// Variant 2 ( can consume 2 args )
	val := os.Args[0]
	// convert val to boolean ( https://golang.org/pkg/strconv/ )
	_, err := strconv.ParseBool(val)
	// if conditions can only work with boolean
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(os.Args[0] + " " + os.Args[1])

}