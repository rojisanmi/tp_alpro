package main

import "fmt"

func printPolaSegitiga(baris int, kolom int, n int, angka int) {
	/* I.S Terdefinisi nilai bilangan bulat baris, kolom, dan n
	   F.S menampilkan pola string "*" yang berbentuk segitiga */
	if kolom <= n {
		if baris+kolom <= n {
			fmt.Print(" ")
		} else {
			fmt.Print(angka % 10)
			angka = angka + 1
		}
		printPolaSegitiga(baris, kolom+1, n, angka) // gunakan fungsi rekursif pada baris ini
	} else if baris < n {
		fmt.Println()
		printPolaSegitiga(baris+1, 1, n, angka) // gunakan fungsi rekursif pada baris ini
	}
}

func main() {
	var n, angka int
	fmt.Scan(&n, &angka)
	printPolaSegitiga(1, 1, n, angka)
}
