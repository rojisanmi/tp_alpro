package main

import "fmt"

type tabInt [5]int

func main() {
	var A tabInt
	var N int
	A = tabInt{5, 7, 9, 11, 13}
	N = 5
	mengurutkan(&A, N)
	for i := 0; i < N; i++ {
		fmt.Print(A[i], " ")
	}
}

func mengurutkan(A *tabInt, N int) {
	var pass, idx, i, temp int
	pass = 1
	for pass <= N-1 {
		idx = pass - 1
		i = pass
		for i < N {
			if A[idx] < A[i] {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}
