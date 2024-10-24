package main

import (
	"fmt"
)

type Pasukan struct {
	Nama     string
	Artileri int
	Kavaleri int
}

func main() {
	// Membaca input data pasukan
	var pasukan []Pasukan
	for i := 0; i < 3; i++ {
		var nama string
		var artileri, kavaleri int
		fmt.Scan(&nama, &artileri, &kavaleri)
		pasukan = append(pasukan, Pasukan{Nama: nama, Artileri: artileri, Kavaleri: kavaleri})
	}

	// Mencari pasukan dengan total kekuatan terbanyak
	var pasukanTerkuat Pasukan
	for _, pasukan := range pasukan {
		totalKekuatan := pasukan.Artileri + pasukan.Kavaleri
		if totalKekuatan > pasukanTerkuat.Artileri+pasukanTerkuat.Kavaleri {
			pasukanTerkuat = pasukan
		}
	}

	// Mencetak nama pasukan dengan total kekuatan terbanyak
	fmt.Print(pasukanTerkuat.Nama)
}
