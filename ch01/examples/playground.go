package main

import (
	"runtime"
	"fmt"
)

func main(){

	version := runtime.Version()
	fmt.Printf("Go Version: %s\n", version)

}