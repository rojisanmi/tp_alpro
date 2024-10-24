package main

import (
	"fmt"
	"math"
)

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var data tabInt
	var ndata int
	// baca data
	baca(&data, &ndata)
	// panggil dan cetak fungsi standardDeviation
	fmt.Print(standardDeviation(data, ndata))
}

func baca(A *tabInt, n *int) {
	/*
		IS: Array (A) dan banyak elemen (n) terdefinisi sembarang
		FS: Array (A) berisi bilangan bulat sejumlah n elemen
	*/
	fmt.Scan(&*n)
	for i := 0; i < *n; i++ {
	    fmt.Scan(&(*A)[i])
	}
}

func rerata(A tabInt, n int) float64 {
	/* mengembalikan nilai rata-rata n elemen Array (A) */
	var jumlah float64
	var jumlahint int
    for i := 0; i < n; i++ {
        jumlahint += A[i]
    }
    jumlah = float64(jumlahint)
    return jumlah / float64(n)
}

func standardDeviation(A tabInt, n int) float64 {
	/* Mengembalikan nilai Simpangan Baku untuk sampel n elemen */
	var hasil, avg float64
	avg = rerata(A, n)
	for i := 0; i < n; i++ {
	    hasil += math.Pow((float64(A[i])-avg), 2)/(float64(n) - 1)
	}
	return math.Sqrt(hasil)
}
