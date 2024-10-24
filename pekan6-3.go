package main

import "fmt"

type order struct {
	nomorMeja  string
	makanan    string
	minuman    string
	totalPrice int
}

func main() {
	var order1, order2 order
	fmt.Scanln(&order1.nomorMeja, &order1.makanan, &order1.minuman)
	fmt.Scanln(&order2.nomorMeja, &order2.makanan, &order2.minuman)
	orderPrioritas(order1, order2)
}

func totalPesanan(menu order) int {
	//I.S terdefinisi order
	//F.S Mengembalikan total pesanan dari sebuah order
	var harga1, harga2 int
	if menu.makanan == "Dimsum" {
		harga1 = 20000
	} else if menu.makanan == "Bakwan" {
		harga1 = 12000
	} else if menu.makanan == "Es_Teh" {
		harga1 = 7000
	} else if menu.makanan == "Air" {
		harga1 = 3000
	}

	if menu.minuman == "Dimsum" {
		harga2 = 20000
	} else if menu.minuman == "Bakwan" {
		harga2 = 12000
	} else if menu.minuman == "Es_Teh" {
		harga2 = 7000
	} else if menu.minuman == "Air" {
		harga2 = 3000
	}

	return harga1 + harga2
}

func orderPrioritas(order1, order2 order) {
	//I.S terdefinisi order1, dan order2
	//F.S Mengembalikan nomor meja yang lebih dulu diprioritaskan
	// Jika prioritas sama tampilkan keduanya
	order1.totalPrice = totalPesanan(order1)
	order2.totalPrice = totalPesanan(order2)

	if order1.totalPrice > order2.totalPrice {
		fmt.Println("Nomor Meja:", order1.nomorMeja)
		fmt.Print("Total order hari ini: Rp ", order1.totalPrice+order2.totalPrice)
	} else if order1.totalPrice < order2.totalPrice {
		fmt.Println("Nomor Meja:", order2.nomorMeja)
		fmt.Print("Total order hari ini: Rp ", order1.totalPrice+order2.totalPrice)
	} else {
		fmt.Println("Nomor Meja:", order1.nomorMeja)
		fmt.Println("Nomor Meja:", order2.nomorMeja)
		fmt.Print("Total order hari ini: Rp ", order1.totalPrice+order2.totalPrice)
	}
}
