package main

import "fmt"

type pegawai struct {
	nama                    string
	gaji_pokok, besar_bonus int
	masa_kerja              float64
}

func main() {
	var staff pegawai

	fmt.Scan(&staff.nama, &staff.gaji_pokok, &staff.masa_kerja)

	hitung_bonus(&staff)

	fmt.Printf("Pegawai dengan nama %s mendapatkan bonus sebesar Rp %d", staff.nama, staff.besar_bonus)
}

func hitung_bonus(p *pegawai) {
	if *&p.masa_kerja >= 10 {
		*&p.besar_bonus = *&p.gaji_pokok * 3 / 2
	} else if *&p.masa_kerja >= 5 {
		*&p.besar_bonus = *&p.gaji_pokok * 3 / 4
	} else {
		*&p.besar_bonus = *&p.gaji_pokok / 2
	}
}
