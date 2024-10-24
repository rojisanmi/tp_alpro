package main

import "fmt"

func main() {
	var pilihan int

	for {
		_, err := fmt.Scanf("%d", &pilihan)
		if err != nil {
			continue
		}
		menu()
		switch pilihan {
		case 1:
			hitungJumlah()
		case 2:
			hitungKali()
		case 3:
			hitungBagi()
		case 4:
			fmt.Println("Pilih (1/2/3/4)? 4")
			break
		}
		if pilihan == 4 {
			break
		}
	}
}

func menu() {
	fmt.Println("-----------------------")
	fmt.Printf("%15s\n", "M E N U")
	fmt.Println("-----------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Println("-----------------------")
}

func hitungJumlah() {
	var bil1, bil2, hasil int
	fmt.Println("Pilih (1/2/3/4)? 1")
	fmt.Scan(&bil1, &bil2)
	fmt.Println("Masukkan dua bilangan yang akan dijumlahkan:", bil1, bil2)
	hasil = bil1 + bil2
	fmt.Println("Hasil penjumlahan:", hasil)
}

func hitungKali() {
	var bil1, bil2, hasil int
	fmt.Println("Pilih (1/2/3/4)? 2")
	fmt.Scan(&bil1, &bil2)
	fmt.Println("Masukkan dua bilangan yang akan dikalikan:", bil1, bil2)
	hasil = bil1 * bil2
	fmt.Println("Hasil perkalian:", hasil)
}

func hitungBagi() {
	var bil1, bil2, hasil float64
	fmt.Println("Pilih (1/2/3/4)? 3")
	fmt.Scan(&bil1, &bil2)
	fmt.Println("Masukkan dua bilangan yang akan dibagikan:", bil1, bil2)
	hasil = bil1 / bil2
	fmt.Printf("Hasil pembagian: %.1f\n", hasil)
}
