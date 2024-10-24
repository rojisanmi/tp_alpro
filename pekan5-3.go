package main

import "fmt"

func perkalianBilanganGenap(n int) int {
	/* Mengembalikan sebuah bilangan bulat yang menyatakan
	   hasil perkalian bilangan genap dari 2 hingga N */
	if n < 2 { // basis: jika n kurang dari 2, kembalikan 1
		return 1
	} else if n%2 != 0 { // jika n ganjil, kurangi satu agar menjadi genap
		n--
	}
	return n * perkalianBilanganGenap(n-2)
}

func main() {
	// Contoh penggunaan fungsi perkalianBilanganGenap
	fmt.Println("Input: 2, Output:", perkalianBilanganGenap(2))   // 2
	fmt.Println("Input: 3, Output:", perkalianBilanganGenap(3))   // 2
	fmt.Println("Input: 4, Output:", perkalianBilanganGenap(4))   // 8
	fmt.Println("Input: 5, Output:", perkalianBilanganGenap(5))   // 8
	fmt.Println("Input: 10, Output:", perkalianBilanganGenap(10)) // 3840
}
