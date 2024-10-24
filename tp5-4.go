package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Masukkan nilai n: ")
	fmt.Scanln(&n)

	// Menghitung nilai suku ke-n
	nilaiSuku := nilaiSukuRekursif(n)
	fmt.Println("Nilai suku ke-", n, ":", nilaiSuku)
}

func nilaiSukuRekursif(n int) int {
	if n == 1 {
		return 1
	} else if n <= 3 {
		return nilaiSukuRekursif(n - 1)
	} else {
		return nilaiSukuRekursif(n-1) + nilaiSukuRekursif(n-2) + nilaiSukuRekursif(n-3)
	}
}
