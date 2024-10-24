package main

import "fmt"

func main() {
	type topskor struct {
		nama   string
		gol    float64
		assist float64
	}
	var jumlah_gol, jumlah_assist, rerata_gol, rerata_assist float64
	var stats topskor

	fmt.Scan(&stats.nama, &stats.gol, &stats.assist)
	jumlah_gol += stats.gol
	jumlah_assist += stats.assist
	fmt.Scan(&stats.nama, &stats.gol, &stats.assist)
	jumlah_gol += stats.gol
	jumlah_assist += stats.assist
	fmt.Scan(&stats.nama, &stats.gol, &stats.assist)
	jumlah_gol += stats.gol
	jumlah_assist += stats.assist

	rerata_gol = jumlah_gol / 3
	rerata_assist = jumlah_assist / 3
	fmt.Println(rerata_gol)
	fmt.Println(rerata_assist)
}
