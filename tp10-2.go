package main

import "fmt"

const NMAX int = 11

type pemain struct {
	nama, nomorPunggung, posisi string
	tinggi                      int
}

// tabPemain adalah alias array pemain dengan maks elemen NMAX
type tabPemain [NMAX]pemain

func main() {
	var timnas tabPemain
	var nPemain int
	nPemain = 0
	// baca data
	bacaData(&timnas, &nPemain)
	// cetak data pemain
	cetakPemain(timnas, nPemain)
	cetakNamaPemainTertinggi(timnas, nPemain)
	// cetak nama pemain terendah
	cetakNamaPemainTerendah(timnas, nPemain)
}

func bacaData(A *tabPemain, n *int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi sembarang
		Proses: Membaca nama, nomor punggung, posisi, dan tinggi badan
			 	Jika array belum penuh dan nama bukan "none", maka nama, nomor punggung, posisi,
				dan tinggi badan dimasukkan ke dalam array.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	var name, np, p string
	var t int
	var i int

	i = 0
	fmt.Scanln(&name, &np, &p, &t)
	for name != "none" && i < NMAX {
		A[i].nama = name
		A[i].nomorPunggung = np
		A[i].posisi = p
		A[i].tinggi = t
		i++
		fmt.Scanln(&name, &np, &p, &t)
	}
	*n = i
	if *n > NMAX {
		*n = NMAX
	}
}

func cetakPemain(A tabPemain, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Data pemain:
			<nama1> <nomorPunggung1> <posisi1> <tinggi1>
			<nama2> <nomorPunggung2> <posisi2> <tinggi2>
			...
			<naman> <nomorPunggungn> <posisin> <tinggin>"
	*/
	var i int
	fmt.Println("Data Pemain:")
	for i = 0; i < n; i++ {
		fmt.Printf("%s %s %s %d\n", A[i].nama, A[i].nomorPunggung, A[i].posisi, A[i].tinggi)
	}
}

func cetakNamaPemainTertinggi(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan tertinggi dengan format:
	       "Pemain dengan badan tertingggi: <nama>"
		   Asumsi: Hanya ada 1 pemain dengan badan tertinggi
	*/
	var nama_tertinggi string
	var tinggi_max int
	var i int
	tinggi_max = A[0].tinggi
	for i = 1; i < n; i++ {
		if A[i].tinggi > tinggi_max {
			nama_tertinggi = A[i].nama
			tinggi_max = A[i].tinggi
		}
	}
	fmt.Printf("Pemain dengan badan tertinggi: %s\n", nama_tertinggi)
}

func cetakNamaPemainTerendah(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan terendah dengan format:
	       "Pemain dengan badan terendah: <nama>""
		   Asumsi: Hanya ada 1 pemain dengan badan terendah
	*/
	var nama_terendah string
	var tinggi_min int
	var i int
	tinggi_min = A[0].tinggi
	for i = 1; i < n; i++ {
		if A[i].tinggi < tinggi_min {
			nama_terendah = A[i].nama
			tinggi_min = A[i].tinggi
		}
	}
	fmt.Printf("Pemain dengan badan terendah: %s\n", nama_terendah)
}
