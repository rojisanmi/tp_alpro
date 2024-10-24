package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(biner(x))
}

func biner(x int) int {
	if x <= 0 {
		return 0
	} else {
		return x%2 + 10*biner(x/2)
	}
}