package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/serkansipahi/gopl-exercise-tasks/ch02/ex.2.1/tempconv"
)


func main(){

	arg := []string{"foo", "33"}

	if len(os.Args) > 1 {
		arg = os.Args
	}

	for _, arg := range arg[1:] {

		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f),
			c, tempconv.CToF(c),
		)
	}

}

