package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	var paket tabBarang
	fmt.Scan(&N)
	isiArray(&paket, N)
	isiBiaya(&paket, N)
	cetakBiaya(paket, N)
}

type barang struct {
	kodePaket       string
	namaPengirim    string
	tujuan          string
	beratBarang     float64
	biayaPengiriman int
}
type tabBarang [1000]barang

func isiArray(arrbarang *tabBarang, n int) {
	/*I.S. Terdefinisi nilai bilangan bulat n. Data barang yang akan dikirim dan dihitung biaya pengirimannya tersedia
	  pada input device sebanyak n*/
	/*F.S. Array arrbarang telah berisi data yang diberikan*/
	for i := 0; i < n; i++ {
		fmt.Scan(&arrbarang[i].kodePaket, &arrbarang[i].namaPengirim, &arrbarang[i].tujuan, &arrbarang[i].beratBarang)
	}
}

func biaya(x barang) int {
	/*Mengembalikan biaya dari paket sesuai dengan ketentuan yang berlaku */
	if x.tujuan == "Jakarta" {
		if x.beratBarang > 40 {
			return 800 + (1000 * 40) + (((int(math.Ceil(x.beratBarang))) % 40) * 4000)
		} else if x.beratBarang > 20 {
			return 800 + (1000 * 20) + (((int(math.Ceil(x.beratBarang))) % 20) * 2000)
		} else {
			return 800 + (1000 * (int(math.Ceil(x.beratBarang))))
		}
	} else {
		if x.beratBarang > 40 {
			return 800 + (1500 * 40) + (((int(math.Ceil(x.beratBarang))) % 40) * 4000)
		} else if x.beratBarang > 20 {
			return 800 + (1500 * 20) + (((int(math.Ceil(x.beratBarang))) % 20) * 2000)
		} else {
			return 800 + (1500 * (int(math.Ceil(x.beratBarang))))
		}
	}
}
func isiBiaya(arrbarang *tabBarang, n int) {
	/*I.S. Terdefinisi nilai bilangan bulat n. Data barang yang akan dikirim dan dihitung biaya pengirimannya tersedia
	  pada input device sebanyak n*/
	/*F.S. nilai biaya pengiriman pada array telah berisi data*/
	for i := 0; i < n; i++ {
		arrbarang[i].biayaPengiriman = biaya(arrbarang[i])
	}
}

func cetakBiaya(arrbarang tabBarang, n int) {
	/*I.S. Terdefinisi array arrBarang yang berisi sebanyak n data barang yang akan dikirim beserta biaya pengiriman disetiap paket
	  F.S. Menampilkan ke layar daftar kode paket, nama pengirim dan biaya pengiriman*/
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Printf("%s %s %d", arrbarang[i].kodePaket, arrbarang[i].tujuan, arrbarang[i].biayaPengiriman)
		} else {
			fmt.Printf("%s %s %d\n", arrbarang[i].kodePaket, arrbarang[i].tujuan, arrbarang[i].biayaPengiriman)
		}
	}
}
