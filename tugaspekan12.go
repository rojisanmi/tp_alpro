package main

import "fmt"

// Definisi konstanta dan tipe data sesuai dengan kamus
const nMAX int = 2022

type student struct {
	name string
	sid  string
	gpa  float64
}

type tabMhs [nMAX]student

// Procedure untuk mengurutkan data mahasiswa berdasarkan IPK secara ascending menggunakan selection sort
func mengurutkan(arr *tabMhs, n int) {
	var pass, idx, i int
	var temp float64
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if arr[i].gpa < arr[idx].gpa {
				idx = i
			}
			i = i + 1
		}
		temp = arr[pass-1].gpa
		arr[pass-1] = arr[idx]
		arr[idx].gpa = temp
		pass = pass + 1
	}
}

// Fungsi utama untuk menjalankan program
func main() {
	var students tabMhs
	var n int

	// Memasukkan data mahasiswa (contoh data)
	students[0] = student{"Alice", "A001", 3.5}
	students[1] = student{"Bob", "B002", 3.8}
	students[2] = student{"Charlie", "C003", 3.2}
	students[3] = student{"David", "D004", 3.6}
	n = 4 // Jumlah data mahasiswa yang dimasukkan

	// Menampilkan data mahasiswa sebelum diurutkan
	fmt.Println("Data mahasiswa sebelum diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("Name: %s, SID: %s, GPA: %.2f\n", students[i].name, students[i].sid, students[i].gpa)
	}

	// Mengurutkan data mahasiswa menggunakan selection sort
	mengurutkan(&students, n)

	// Menampilkan data mahasiswa setelah diurutkan
	fmt.Println("\nData mahasiswa setelah diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("Name: %s, SID: %s, GPA: %.2f\n", students[i].name, students[i].sid, students[i].gpa)
	}
}
