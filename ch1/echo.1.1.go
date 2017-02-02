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

	// Variant 2 ( can consue 2 args )
	val := os.Args[0];
	// convert val to boolean
	_, err := strconv.ParseBool(val)
	// conditions can only work with boolean
	if err == nil {
		fmt.Println(os.Args[0] + " " + os.Args[1])
	}


}