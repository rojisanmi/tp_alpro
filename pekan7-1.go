package main

import "fmt"

const NMax = 1000

func jumlahBilangan(bilangan [NMax]int) int {
	var jumlah int
	for _, n := range bilangan {
		if n == 0 {
			break
		}
		jumlah += n
	}
	return jumlah
}

func main() {
	var bilangan [NMax]int
	var input int
	fmt.Scan(&input)
	i := 0
	for input != 0 {
		bilangan[i] = input
		i++
		fmt.Scan(&input)
	}

	jumlah := jumlahBilangan(bilangan)
	fmt.Print(jumlah)
}
