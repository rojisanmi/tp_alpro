package main

import "fmt"

func capitalize(s string, n int) {
	if n > 0 {
		capitalize(s, n-1)
		fmt.Print(string(s[n-1] - 32))
	}
}

func main() {
	var s string
	fmt.Println("Masukkan string: ")
	fmt.Scanln(&s)

	capitalize(s, len(s))
}
