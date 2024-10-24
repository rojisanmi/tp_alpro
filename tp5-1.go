package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci_rekursif(n))
}

func jumlah_bilangan_asli_rekursif(bil int) int {
	if bil == 0 {
		return 0
	} else {
		return bil + jumlah_bilangan_asli_rekursif(bil-1)
	}
}

func faktorial_rekursif(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * faktorial_rekursif(n-1)
	}
}

func fibonacci_rekursif(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 || n == 3 {
		return 1
	} else {
		return fibonacci_rekursif(n-1) + fibonacci_rekursif(n-2)
	}
}
