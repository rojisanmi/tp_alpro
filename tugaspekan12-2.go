package main

import "fmt"

const MAX int = 30

type IntArray struct {
	tabInt [MAX + 1]int
	N      int
}

var array1, array2 IntArray

func main() {
	inputArray(&array1)
	inputArray(&array2)
	sortArray(&array1)
	sortArray(&array2)
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(isSimilar(array1, array2))
}

func inputArray(T *IntArray) {
	var input int
	T.N = 1
	fmt.Println("Masukkan elemen array (akhiri dengan 0):")
	for {
		fmt.Scan(&input)
		if input == 0 || T.N >= MAX {
			break
		}
		T.tabInt[T.N] = input
		T.N++
	}
}

func sortArray(T *IntArray) {
	var pass, idx, i, temp int
	pass = 1
	for pass <= T.N {
		idx = pass
		i = pass + 1
		for i <= T.N {
			if T.tabInt[i] < T.tabInt[idx] {
				idx = i
			}
			i = i + 1
		}
		temp = T.tabInt[pass-1]
		T.tabInt[pass-1] = T.tabInt[idx]
		T.tabInt[idx] = temp
		pass = pass + 1
	}
}

func isSimilar(T1, T2 IntArray) bool {
	if T1.N != T2.N {
		return false
	}
	for i := 0; i < T1.N; i++ {
		if T1.tabInt[i] != T2.tabInt[i] {
			return false
		}
	}
	return true
}
