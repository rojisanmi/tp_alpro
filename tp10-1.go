package main

import "fmt"

// Deklarsi konstanta 20 elemen
const NMAX int = 20

// Deklarasi tabInt dengan total NMAX elemen
type tabInt [NMAX]int

func main() {
	// Deklarasi variabel bertipe tabInt
	var data tabInt
	// Deklarasi banyak elemen array
	var n int
	// Pembacaan data bilangan
	baca(&data, &n)
	// Cetak elemen array
	cetakElemen(data, n)
	// Cetak Info
	cetakInfo(data, n)
}

func baca(A *tabInt, n *int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi sembarang
		Proses: Membaca bilangan bulat dan memasukkan bilangan bulat positif
		        ke dalam array A.Pembacaan dilakukan selama terbaca
		        bilangan positif dan n < NMAX.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	var x, i int
	for i < NMAX {
		fmt.Scan(&x)
		if x <= 0 {
			break
		}
		(*A)[i] = x
		i++
	}
	*n = i

}

func cetakElemen(A tabInt, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Elemen array: <e1> <e2> <e3> ... <en>"
	*/
	var i int
	fmt.Print("Elemen array:")
	for i = 0; i < n; i++ {
		fmt.Printf(" %d", A[i])
	}
	fmt.Println()
}

func maksimum(A tabInt, n int) int {
	/* Mengembalikan nilai elemen maksimum dari array A dengan banyak elemen n */
	var i, max int
	max = A[0]
	for i = 1; i < n; i++ {
		if A[i] > max {
			max = A[i]
		}
	}
	return max
}

func minimum(A tabInt, n int) int {
	/* Mengembalikan nilai elemen minimum dari array A dengan banyak elemen n */
	var i, min int
	min = A[0]
	for i = 1; i < n; i++ {
		if A[i] < min {
			min = A[i]
		}
	}
	return min
}

func cetakInfo(A tabInt, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Nilai maksimum, nilai minimum, dan banyaknya elemen tercetak dengan format:
			"Nilai maksimum: <max_value>
			 Nilai minimum: <min_value>
			 Banyak elemen: <n>"
	*/

	fmt.Printf("Nilai maksimum: %d\n", maksimum(A, n))
	fmt.Printf("Nilai minimum: %d\n", minimum(A, n))
	fmt.Printf("Banyak elemen: %d\n", n)

}
