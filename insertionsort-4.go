package main

import "fmt"

const nMax = 1000

type peserta struct {
	nama                 string
	gold, silver, bronze int
}

type olimpiade [nMax]peserta

func main() {
	var olimpiade olimpiade
	var nPeserta int
	olimpiade[0] = peserta{"mpvh", 8, 4, 8}
	olimpiade[1] = peserta{"ptpc", 2, 7, 10}
	olimpiade[2] = peserta{"omen", 8, 9, 5}
	olimpiade[3] = peserta{"rmpw", 2, 8, 7}
	olimpiade[4] = peserta{"dnba", 7, 8, 1}
	nPeserta = 5
	sorting(&olimpiade, nPeserta)
	tampilArray(olimpiade, nPeserta)
}

func tampilArray(t olimpiade, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Println(t[i].nama, t[i].gold, t[i].silver, t[i].bronze)
	}
}

func poin(g, s, b int) int {
	return (4 * g) + (3 * s) + b
}

func sorting(t *olimpiade, n int) {
	var i, pass int
	var temp peserta
	pass = 1
	for pass < n {
		i = pass
		temp = t[pass]
		for (i > 0 && poin(temp.gold, temp.silver, temp.bronze) > poin(t[i-1].gold, t[i-1].silver, t[i-1].bronze)) ||
			(i > 0 && poin(temp.gold, temp.silver, temp.bronze) == poin(t[i-1].gold, t[i-1].silver, t[i-1].bronze) &&
				temp.nama[0] > t[i-1].nama[0]) {
			t[i] = t[i-1]
			i--
		}
		t[i] = temp
		pass++
	}
}
