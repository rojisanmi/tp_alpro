package main

import "fmt"

func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm += 1
	}
}

func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd += 1
	}
}

func hitungKalah(g, k int, jk *int) {
	if g < k {
		*jk += 1
	}
}

func hitungJumGolKegolanSelisih(g, k int, jg, jk, jsg *int) {
	*jg += g
	*jk += k
	*jsg += g - k
}

func hitungJumPoint(jp *int) {
	*jp = *jp * 3
}

func main() {
	var m, g, k, jm, jd, jk, jg, jkg, jsg, jp int
	fmt.Scan(&m)
	for i := 1; i <= m; i++ {
		fmt.Scan(&g, &k)
		hitungMenang(g, k, &jm)
		hitungDraw(g, k, &jd)
		hitungKalah(g, k, &jk)
		hitungJumGolKegolanSelisih(g, k, &jg, &jkg, &jsg)
	}
	jp = jm
	hitungJumPoint(&jp)
	jp += jd
	fmt.Print(m, jm, jd, jk, jg, jkg, jsg, jp)
}
