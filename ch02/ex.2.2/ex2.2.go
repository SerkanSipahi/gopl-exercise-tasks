package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/serkansipahi/gopl-exercise-tasks/ch02/ex.2.1/tempconv"
	"bufio"
)

func process(arg string){

	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	{
		feet := tempconv.Feet(t)
		meter := tempconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n",
			feet, tempconv.FeetToMeter(feet),
			meter, tempconv.MeterToFeet(meter),
		)
	}


	{
		pound := tempconv.Pound(t)
		kilogram := tempconv.Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n",
			pound, tempconv.PoundToKilogram(pound),
			kilogram, tempconv.KilogramToPound(kilogram),
		)
	}


	{
		celsius := tempconv.Celsius(t)
		fahrenheit := tempconv.Fahrenheit(t)
		fmt.Printf("%s = %s, %s = %s\n",
			celsius, tempconv.CToF(celsius),
			fahrenheit, tempconv.FToC(fahrenheit),
		)
	}
}

func main(){

	if len(os.Args) > 1 {
		for _, value := range os.Args[1:] {
			process(value)
		}
		return
	}

	fmt.Println("Type and apply with CTRL + D")
	fmt.Println("With CTRL + C quit programm")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		process(scanner.Text())
	}

}

