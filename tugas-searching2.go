package main

import "fmt"

const nMax int = 100

type himpunan [nMax]int

func main() {
	var A, B, irisan himpunan
	A = himpunan{11, 28, 33, 64, 95, 16, 100, 15}
	B = himpunan{3, 11, 7, 28, 33, 6}
	irisan = mengiris(A, B)
	fmt.Print(irisan)
}

func mengiris(A, B himpunan) himpunan {
	var irisan himpunan
	var i, j int
	j = 0
	for i = 0; i < nMax; i++ {
		if A[i] != 0 && ada(B, A[i]) {
			irisan[j] = A[i]
			j++
		}
	}
	return irisan
}

func ada(B himpunan, angka int) bool {
	for i := 0; i < nMax; i++ {
		if B[i] == angka {
			return true
		}
	}
	return false
}
