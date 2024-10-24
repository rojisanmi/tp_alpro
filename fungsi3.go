package main

import "fmt"

// implementasikan fungsinya saja

func jumlahGenap(awal, akhir int) int {
	var total int
	for i := awal; i <= akhir; i++ {
		if i%2 == 0 {
			total += i
		}
	}
	return total
}

func main() {
	var bil1, bil2 int
	fmt.Scan(&bil1, &bil2)
	fmt.Println(jumlahGenap(bil1, bil2))
}
