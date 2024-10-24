package main

import "fmt"

func sumCubic(n int) int {
	if n == 1 {
		return 1
	} else {
		return n*n*n + sumCubic(n-1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(sumCubic(n))
}
