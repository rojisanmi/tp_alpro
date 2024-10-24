package main

import "fmt"

const NMAX int = 20

type tabChar [NMAX + 1]byte

func main() {
	var kata tabChar
	var nChar int
	fmt.Scan(&nChar)
	baca(&kata, &nChar)
	cetak(kata, nChar)
}

func baca(k *tabChar, n *int) {
	for i := 0; i <= *n; i++ {
		if i <= 20 {
			fmt.Scanf("%c", &k[i])
		}
	}
}

func cetak(k tabChar, n int) {
	for i := n; i >= 0; i -= 1 {
		if i <= 20 && i > 0 {
			fmt.Printf("%c", k[i])
		}
	}
}
