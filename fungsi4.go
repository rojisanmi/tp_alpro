package main

import "fmt"

// implementasikan fungsinya saja

func hitungPerkalian(b1, b2 int) int {
	/* function mengembalikan hasil perkalian 2 bilangan, jika kedua bilangan genap
	   kembalikan bilangan 0 jika keduanya tidak genap
	*/
	var result int

	if b1%2 == 1 || b2%2 == 1 {
		b1 = 0
		b2 = 0
	}
	result = b1 * b2
	return result
}

func main() {
	var b1, b2, hasil int
	fmt.Scan(&b1, &b2)
	hasil = hitungPerkalian(b1, b2)
	fmt.Println(hasil)
}
