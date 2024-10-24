package main

import (
	"fmt"
	"math"
)

const NMAX int = 10

type tabInt [NMAX + 1]int

func main() {
	var data tabInt
	var nData int
	baca(&data, &nData)
	if nData > 10 {
		nData = 10
	}
	cetak(data, nData)
	fmt.Print(jumlah(data, nData), rata_rata(data, nData))
}

func baca(A *tabInt, n *int) {
	var angka int
	i := 0
	fmt.Scan(&A[i])
	for A[i] != 0 {
		*n += 1
		i += 1
		fmt.Scan(&angka)
		if *n < 10 {
			A[i] = angka
		}
	}
}

func cetak(A tabInt, n int) {
	if n == 0 {
		fmt.Println()
	} else {
		for i := 0; i < n; i++ {
			if A[i] > 0 {
				fmt.Print(A[i], " ")
			}
		}
		fmt.Println()
	}
}

func jumlah(A tabInt, n int) int {
	var angka int
	sum := 0
	for i := 0; i < n; i++ {
		if A[i] < 0 {
			angka = A[i] * -1
		} else {
			angka = A[i]
		}
		sum += angka
	}
	return sum
}

func rata_rata(A tabInt, n int) float64 {
	var angka int
	if n == 0 {
		result := math.NaN()
		return result
	} else {
		sum := 0
		for i := 0; i < n; i++ {
			if A[i] < 0 {
				angka = A[i] * -1
			} else {
				angka = A[i]
			}
			sum += angka
		}
		rerata := float64(sum) / float64(n)
		return rerata
	}
}
