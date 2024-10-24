package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	cetakRerata(n, 1, 0, 0)
}

func cetakRerata(n, i, jumlah int, rata float64) {
	if i > n {
		return
	} else {
		jumlah += i
		rata = float64(jumlah) / float64(i)
		fmt.Println(jumlah, rata)
		cetakRerata(n, i+1, jumlah, rata)
	}
}
