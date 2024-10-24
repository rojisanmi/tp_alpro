package main

import "fmt"

const nMax int = 100

type himpunan [nMax]int

func main() {
	var A, B, gabungan himpunan
	A = himpunan{11, 28, 33, 64, 95, 16, 100, 15}
	B = himpunan{3, 11, 7, 28, 33, 6}
	gabungan = menggabungkan(A, B)
	fmt.Print(gabungan)
}

func menggabungkan(A, B himpunan) himpunan {
	var gabungan himpunan
	var i, j int
	gabungan = A
	j = panjangHimpunan(A) - 1
	for i = 0; i < nMax; i++ {
		if B[i] != 0 && ada(A, B[i]) != true {
			gabungan[j] = B[i]
			j++
		}
	}
	return gabungan
}

func ada(B himpunan, angka int) bool {
	for i := 0; i < nMax; i++ {
		if B[i] == angka {
			return true
		}
	}
	return false
}

func panjangHimpunan(A himpunan) int {
	var panjang, i int
	panjang = 0
	for i = 0; i < nMax; i++ {
		if A[i] != 0 {
			panjang++
		}
	}
	return panjang
}
