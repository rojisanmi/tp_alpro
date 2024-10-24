package main

import "fmt"

const NMAX int = 5

type tabInt struct {
	info [NMAX]int
	n    int
}

func main() {
	var nilaiAkhir tabInt
	nilaiAkhir.n = 1
	for i := 0; i < 5; i++ {
		if i <= 4 {
			fmt.Scan(&nilaiAkhir.info[i])
			nilaiAkhir.n += 1
		}
	}
	cetakNilai(nilaiAkhir)
}

func cetakNilai(NA tabInt) {
	for i := 0; i < NA.n; i++ {
		if i <= 4 {
			fmt.Print(NA.info[i], " ")
		}
	}
}
