package main

import "fmt"

const nMax int = 37

type tHimpunan struct {
	anggota [nMax]int
	panjang int
}

func main() {
	var himpunan1, himpunan2 tHimpunan
	himpunan1.anggota = [37]int{8, 2, 4, 1, 5, 4}
	himpunan1.panjang = 6
	himpunan2.anggota = [37]int{8, 2, 4, 1, 5, 1}
	himpunan2.panjang = 6
	fmt.Print("Himpunan 1 = Himpunan 2? ")
	fmt.Println(sama(himpunan1, himpunan2))
}

func ada(set tHimpunan, x int) bool {
	var i int
	var hasil bool
	hasil = false
	i = 0
	for i = 0; i < set.panjang; i++ {
		if set.anggota[i] == x {
			hasil = true
		}
	}
	return hasil
}

func urut(set tHimpunan) {
	var i, pass, temp int
	pass = 1
	for pass < set.panjang {
		i = pass
		temp = set.anggota[pass]
		for i > 0 && temp < set.anggota[i-1] {
			set.anggota[i] = set.anggota[i-1]
			i = i - 1
		}
		set.anggota[i] = temp
		pass = pass + 1
	}
}

func sama(set1, set2 tHimpunan) bool {
	var i int
	urut(set1)
	urut(set2)
	for i = 0; i < set1.panjang; i++ {
		if ada(set2, set1.anggota[i]) == false {
			return false
		}
	}
	return true
}
