package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func distance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}

func main() {
	var a, b, c Point

	fmt.Scan(&a.X, &a.Y)
	fmt.Scan(&b.X, &b.Y)
	fmt.Scan(&c.X, &c.Y)

	perimeter := distance(a, b) + distance(b, c) + distance(c, a)

	fmt.Printf("%.15f", perimeter)
}
