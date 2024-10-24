package main

import "fmt"

const nMax = 1000

type tabString [nMax]string

func main() {
	var tabBunga tabString
	var nBunga int
	tabBunga = tabString{"Mawar.", "Lili.", "Kertas.", "Forget_me_not.", "Kamboja.", "Anggrek.", "Rafflesia."}
	nBunga = 7
	mengurutkan(&tabBunga, nBunga)
	tampilArray(tabBunga, nBunga)
}

func panjang(bunga string) int {
	var panjang, i int
	for i = 0; i < len(bunga); i++ {
		if bunga[i] != '.' && bunga[i] != '_' {
			panjang++
		}
	}
	return panjang
}

func mengurutkan(tabBunga *tabString, n int) {
	var pass, i int
	var temp string
	pass = 1
	for pass < n {
		i = pass
		temp = tabBunga[pass]
		for i > 0 && panjang(temp) < panjang(tabBunga[i-1]) {
			tabBunga[i] = tabBunga[i-1]
			i--
		}
		tabBunga[i] = temp
		pass++
	}
}

func isiArray(tabBunga *tabString, n *int) {
	fmt.Scanln(n)
	if *n > 1000 {
		*n = nMax
	}

	for i := 0; i < *n; i++ {
		fmt.Scanln(&tabBunga[i])
	}
}

func tampilArray(tabBunga tabString, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(tabBunga[i])
	}
}
