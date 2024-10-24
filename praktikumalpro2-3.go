package main

import (
	"fmt"
	"math"
)

func panjangX(R float64, T float64) float64 {
	// Convert the T from degrees to radians
	TInRadians := T * math.Pi / 180

	// Calculate the x coordinate using the formula x = r * cos(theta)
	x := R * math.Cos(TInRadians)

	// Return the x coordinate
	return x
}

func panjangY(R float64, T float64) float64 {
	// Convert the T from degrees to radians
	TInRadians := T * math.Pi / 180

	// Calculate the y coordinate using the formula y = r * sin(theta)
	y := R * math.Sin(TInRadians)

	// Return the y coordinate
	return y
}

func main() {
	var R float64
	var T float64

	fmt.Println("Enter the R: ")
	fmt.Scanf("%f", &R)

	fmt.Println("Enter the T (in degrees): ")
	fmt.Scanf("%f", &T)

	x := panjangX(R, T)
	y := panjangY(R, T)

	fmt.Printf("The length of the R in the x-coordinate is: %.2f\n", x)
	fmt.Printf("The length of the R in the y-coordinate is: %.2f\n", y)
}
