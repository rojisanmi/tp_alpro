package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var nilaiAkhir tabInt
	var banyakNilai int
	fmt.Scan(&banyakNilai)
	for i := 0; i < banyakNilai; i++ {
		if i <= 9 {
			fmt.Scan(&nilaiAkhir[i])
		}
	}
	cetakNilai(nilaiAkhir, banyakNilai)
}

func cetakNilai(NA tabInt, n int) {
	if n == 0 {
		fmt.Print("Array kosong")
	} else {
		fmt.Print("Nilai yang terdapat pada array: ")
		for i := 0; i < n; i++ {
			if i <= 9 {
				fmt.Print(NA[i], " ")
			}
		}
	}
}
