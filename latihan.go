package main

import (
	"fmt"
)

func reamur(celsius float64) float64 {
	return celsius * 4.0 / 5.0
}

func fahrenheit(celsius float64) float64 {
	return celsius*9.0/5.0 + 32.0
}

func kelvin(celsius float64) float64 {
	return celsius + 273
}

func main() {
	var start, end, step float64

	fmt.Scanf("%f %f %f", &start, &end, &step)

	fmt.Printf("%10s %10s %10s %10s\n", "Celsius", "Reaumur", "Fahrenheit", "Kelvin")
	for i := start; i <= end; i += step {
		r := reamur(i)
		f := fahrenheit(i)
		k := kelvin(i)
		if i == end {
			fmt.Printf("%10.2f %10.2f %10.2f %10.2f", i, r, f, k)
		} else {
			fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", i, r, f, k)
		}
	}
}
