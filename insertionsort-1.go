package main

import "fmt"

type intArray []int

func main() {
	var A intArray
	A = intArray{1, 5, 2, 4, 3}
	n := 5
	mengurutkan(&A, n)
	median := findMedian(A, n)
	fmt.Print(median)
}

func mengurutkan(A *intArray, n int) {
	var pass, i, temp int
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = (*A)[pass]
		for i > 0 && temp < (*A)[i-1] {
			(*A)[i] = (*A)[i-1]
			i--
		}
		(*A)[i] = temp
		pass++
	}
}

func findMedian(A intArray, n int) float64 {
	var i int
	var median float64
	if n%2 == 1 {
		i = (n + 1) / 2
		median = float64(A[i-1])
	} else {
		i = n / 2
		median = (float64(A[i-1]) + float64(A[i])) / 2
	}
	return median
}
