package main

import "fmt"

const NMAX int = 255

type Array [NMAX]int

func main() {
	var A, B Array
	var n int
	inputArray(&A, &n)
	reverseArray(A, n, &B)
	if ujiPalindromArray(A, B, n) {
		fmt.Print("Array memiliki pola palindrom")
	} else {
		fmt.Print("Array tidak memiliki pola palindrom")
	}
}

func inputArray(A *Array, n *int) {
	var i int
	fmt.Print("Masukkan nilai n (maksimal banyaknya isi array): ")
	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX
	}
	fmt.Print("Masukkan elemen array: ")
	for i = 0; i < *n; i++ {
		fmt.Scan(&A[i])
	}
}

func reverseArray(A Array, n int, B *Array) {
	var i int
	for i = 0; i < n; i++ {
		B[(n-1)-i] = A[i]
	}
}

func ujiPalindromArray(A, B Array, n int) bool {
	var i int
	var kondisi bool
	i = 0
	kondisi = true
	for i < n && kondisi {
		if A[i] != B[i] {
			kondisi = !kondisi
		}
		i++
	}
	return kondisi
}
