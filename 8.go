package main

import "fmt"

const NMAX int = 10

type tim struct {
	gol  [NMAX]int
	totPertandingan, totGol, totKebobolan, totMenang, totKalah, totDraw, totPoint int
}

func main() {
	var timG, timH tim
	bacaHasil(&timG, &timH)
	fmt.Println("Statistik Tim G")
	cetakHasil(timG)
	fmt.Println()
	fmt.Println("\nStatistik Tim H")
	cetakHasil(timH)
}

func bacaHasil(tG, tH *tim) {
	/*
		IS: Tim G (tG) dan tim H (tH) terdefinisi sembarang
		Proses: Membaca skor pertandingan berupa jumlah gol tim G dan tim H.
				Mengisi field-field sesuai skor kedua tim. Field-fieldnya:
				1. total pertandingan
				2. gol setiap pertandingan
				3. total gol
				4. total kebobolan
				5. total point: point 3 untuk menang, point 1 untuk draw
				6. total menang: Menang, jika gol lebih besar dari gol lawan
				7. total draw: Draw, jika gol sama dengan gol lawan
				8. total kalah: Kalah, jika gol lebih kecil dari gol lawan
				Pembacaan skor berakhir jika kedua skor bernilai negatif.
		FS: Data tim G (tG) dan data tim H (tH) berisi nilai
	*/
	var i int
	i = 0
	fmt.Scan(&tG.gol[i], &tH.gol[i])
	for (*tG).gol[i] >= 0 || (*tH).gol[i] >= 0 {
		if i < NMAX {
			tG.totPertandingan += 1
			tH.totPertandingan += 1
			tG.totGol += tG.gol[i]
			tG.totKebobolan += tH.gol[i]
			tH.totGol += tH.gol[i]
			tH.totKebobolan += tG.gol[i]
			if tG.gol[i] == tH.gol[i] {
				tG.totPoint += 1
				tH.totPoint += 1
				tG.totDraw += 1
				tH.totDraw += 1
			} else if tG.gol[i] > tH.gol[i] {
				tG.totPoint += 3
				tG.totMenang += 1
				tH.totKalah += 1
			} else {
				tH.totPoint += 3
				tH.totMenang += 1
				tG.totKalah += 1
			}
			i += 1
			}
		fmt.Scan(&(*tG).gol[i], &(*tH).gol[i])
	}
}

func cetakHasil(t tim) {
	/*
		IS: Data tim (t) terdefinisi
		FS: tercetak di layar statistik pertandingan tim (t) dengan format:

		Total pertandingan: <total pertandingan>
		Gol tiap pertandingan: <gol1 gol2 ... goln>
		Total menang: <total kemenangan>
		Total draw: <total draw>
		Total kalah: <total kalah>
		Total gol: <total gol>
		Total kebobolan: <total kebobolan>
		Total point: <total point>
	*/
	var n int
	n = min(NMAX, t.totPertandingan)
	if t.totPertandingan > 10 {
		t.totPertandingan = 10
	}
	fmt.Println("Total pertandingan:", t.totPertandingan)
	fmt.Print("Gol tiap pertandingan: ")
	for i := 0; i < n; i++ {
		fmt.Print(t.gol[i], " ")
	}
	fmt.Println()
	fmt.Println("Total menang:", t.totMenang)
	fmt.Println("Total draw:", t.totDraw)
	fmt.Println("Total kalah:", t.totKalah)
	fmt.Println("Total gol:", t.totGol)
	fmt.Println("Total kebobolan:", t.totKebobolan)
	fmt.Print("Total point: ", t.totPoint)
}

func min (a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}