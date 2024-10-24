package main

import "fmt"

const nMax int = 20

type tabelFrekuensi [nMax]int

func main() {
	var A tabelFrekuensi
	A = tabelFrekuensi{11, 28, 33, 64, 95, 16}
	fmt.Print(isValid(A))
	fmt.Print(len(A))
}

func isValid(A tabelFrekuensi) bool {
	var ketemu int
	for i := 0; i < nMax; i++ {
		if A[i] != 0 {
			ketemu = 0
			for j := 0; j < nMax; j++ {
				if A[i] == A[j] {
					ketemu++
				}
			}
			if ketemu > 1 {
				return false
			}
		}
	}
	return true
}
