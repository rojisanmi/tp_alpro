package main

import "fmt"

const NMax int = 1000

type tabInt [NMax]int
type tabStr [NMax]string

func main() {
	var n, i, total int
	var nama, ruangan tabStr
	var jam, biaya tabInt

	//input banyaknya komputer yang ingin dimainkan
	fmt.Scan(&n)

	//input nama, lama jam, ruangan sebanyak n kali
	for i = 0; i < n; i++ {
		fmt.Scan(&nama[i], &jam[i], &ruangan[i])

		//total biaya per orang berdasarkan ruangan
		if ruangan[i] == "gaming" {
			biaya[i] = 14000 * jam[i]
		} else if ruangan[i] == "biasa" {
			biaya[i] = 10000 * jam[i]
		}
	}

	//menampilkan biaya yang harus dibayar per orang
	for i = 0; i < n; i++ {
		fmt.Printf("Total yang harus dibayar %s sebesar %d\n", nama[i], biaya[i])
		total += biaya[i]
	}

	//menampilkan total keseluruhan
	fmt.Print("Biaya total sebesar ", total)
}
